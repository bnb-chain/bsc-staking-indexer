package indexer

import (
	"context"
	"crypto/tls"
	"math/big"
	"net"
	"net/http"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/bnb-chain/bsc-mev-sentry/syncutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/bnb-chain/bsc-staking-indexer/contract/stakecredit"
	"github.com/bnb-chain/bsc-staking-indexer/contract/stakehub"
	"github.com/bnb-chain/bsc-staking-indexer/model"
	"github.com/bnb-chain/bsc-staking-indexer/store"
	"github.com/bnb-chain/bsc-staking-indexer/util/log"
)

var (
	RtyNum      = uint(5)
	RtyAttempts = retry.Attempts(RtyNum)
	RtyDelay    = retry.Delay(time.Millisecond * 50)
	RtyErr      = retry.LastErrorOnly(true)

	dialer = &net.Dialer{
		Timeout:   time.Second,
		KeepAlive: 60 * time.Second,
	}

	transport = &http.Transport{
		DialContext:         dialer.DialContext,
		MaxIdleConnsPerHost: 2000,
		MaxConnsPerHost:     2000,
		IdleConnTimeout:     90 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}

	httpClient = &http.Client{
		Timeout:   5 * time.Second,
		Transport: transport,
	}
)

type Indexer interface {
	Start(ctx context.Context)
}

type Config struct {
	RPCURL                    string
	StartBlockNumber          int64
	NumberOfBlocksForFinality int64
	Delegators                []common.Address
}

func New(ctx context.Context, cfg Config, store store.Store) (Indexer, error) {
	client, err := ethclient.DialOptions(ctx, cfg.RPCURL, rpc.WithHTTPClient(httpClient))
	if err != nil {
		log.Errorw("indexer: failed to dial ethclient", "err", err)
		return nil, err
	}

	cursor, err := store.QueryCursor(ctx)
	if err != nil {
		log.Errorw("indexer: failed to query cursor", "err", err)
		return nil, err
	}
	if cfg.StartBlockNumber > cursor {
		cursor = cfg.StartBlockNumber - 1
	}

	stakeHub, err := stakehub.New(client)
	if err != nil {
		log.Errorw("indexer: failed to new stakehub", "err", err)
		return nil, err
	}

	validatorInfos, err := store.QueryValidatorInfos(ctx)
	if err != nil {
		log.Errorw("indexer: failed to query validator infos", "err", err)
		return nil, err
	}

	stakeCredits := make(map[common.Address]*stakecredit.ContractWithInfo)
	for _, v := range validatorInfos {
		stakeCredits[v.Credit], err = stakecredit.New(v.Operator, v.Credit, client)
		if err != nil {
			log.Errorw("indexer: failed to new stakecredit", "err", err)
			return nil, err
		}
	}

	return &indexer{
		cursor:                    cursor,
		NumberOfBlocksForFinality: cfg.NumberOfBlocksForFinality,
		delegators:                cfg.Delegators,
		store:                     store,
		client:                    client,
		stakeHub:                  stakeHub,
		stakeCredits:              stakeCredits,
	}, nil
}

type indexer struct {
	cursor                    int64
	NumberOfBlocksForFinality int64
	delegators                []common.Address

	store        store.Store
	client       *ethclient.Client
	stakeHub     *stakehub.Contract
	stakeCredits map[common.Address]*stakecredit.ContractWithInfo // credit address -> stake credit
}

func (i *indexer) Start(ctx context.Context) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := i.index(ctx)
			if err != nil {
				log.Errorw("indexer: failed to indexing", "err", err)
			}
		}
	}
}

func (i *indexer) index(ctx context.Context) error {
	var finalizedHeader *types.Header

	if err := RtyDo(func() error {
		var er error
		finalizedHeader, er = i.client.FinalizedHeader(ctx, i.NumberOfBlocksForFinality)
		return er
	}); err != nil {
		log.Errorw("indexer: failed to get finalized header", "err", err)
		return err
	}

	var parentHeader *types.Header
	for number := i.cursor + 1; number <= finalizedHeader.Number.Int64(); number++ {
		var header *types.Header
		if number == finalizedHeader.Number.Int64() {
			header = finalizedHeader
		} else {
			if err := RtyDo(func() error {
				var er error
				header, er = i.client.HeaderByNumber(ctx, big.NewInt(number))
				return er
			}); err != nil {
				log.Errorw("indexer: failed to get header by number", "number", number, "err", err)
				return err
			}
		}

		if parentHeader == nil {
			if err := RtyDo(func() error {
				var er error
				parentHeader, er = i.client.HeaderByNumber(ctx, big.NewInt(number-1))
				return er
			}); err != nil {
				log.Errorw("indexer: failed to get header by number", "number", number-1, "err", err)
				return err
			}
		}

		err := i.handleStakeHubLogs(ctx, header)
		if err != nil {
			log.Errorw("indexer: failed to process stakeHub logs", "err", err)
			return err
		}

		err = i.handleStakeCreditLogs(ctx, header)
		if err != nil {
			log.Errorw("indexer: failed to process stakeCredit logs", "err", err)
			return err
		}

		if isBreatheBlock(parentHeader.Time, header.Time) {
			err = i.handleStakeCreditCalls(ctx, header)
			if err != nil {
				log.Errorw("indexer: failed to process validator and delegator", "err", err)
				return err
			}
		}

		err = i.store.SaveCursor(ctx, number)
		if err != nil {
			log.Errorw("indexer: failed to save cursor", "number", number, "err", err)
			return err
		}

		parentHeader = header
	}

	return nil
}

// handleStakeHubLogs
// getLogs for ValidatorInfo, DelegatorTx, SlashEvent
func (i *indexer) handleStakeHubLogs(ctx context.Context, header *types.Header) error {
	var (
		validatorInfos []*model.ValidatorInfo
		delegateTxs    []*model.DelegateTx
		slashEvents    []*model.SlashEvent
	)

	err := syncutils.BatchRun(
		func() error {
			var err error
			validatorInfos, err = i.stakeHub.ParseValidatorInfos(header.Number.Uint64())
			if err != nil {
				log.Errorw("indexer: failed to parse validator created", "err", err)
				return err
			}

			return nil
		},

		func() error {
			var err error
			delegateTxs, err = i.stakeHub.ParseDelegateTxs(header, i.delegators)
			if err != nil {
				log.Errorw("indexer: failed to parse delegated", "err", err)
				return err
			}

			return nil
		},

		func() error {
			var err error
			slashEvents, err = i.stakeHub.ParseSlashEvents(header)
			if err != nil {
				log.Errorw("indexer: failed to parse slash events", "err", err)
				return err
			}

			return nil
		},
	)

	if err != nil {
		log.Errorw("indexer: failed to process stakeHub events", "err", err)
		return err
	}

	err = syncutils.BatchRun(
		func() error { return i.store.SaveValidatorInfos(ctx, validatorInfos) },
		func() error { return i.store.SaveDelegateTxs(ctx, delegateTxs) },
		func() error { return i.store.SaveSlashEvents(ctx, slashEvents) },
	)

	if err != nil {
		log.Errorw("indexer: failed to save stakeHub logs data", "err", err)
		return err
	}

	for _, v := range validatorInfos {
		if i.stakeCredits[v.Credit] == nil {
			credit, err := stakecredit.New(v.Operator, v.Credit, i.client)
			if err != nil {
				log.Errorw("indexer: failed to new stakecredit", "err", err)
			}
			i.stakeCredits[v.Credit] = credit
		}
	}

	return nil
}

// handleStakeCreditLogs
// getLogs for BreathBlockRewardEvent
func (i *indexer) handleStakeCreditLogs(ctx context.Context, header *types.Header) error {
	var events []*model.BreathBlockRewardEvent
	for _, v := range i.stakeCredits {
		breatheEvents, err := v.ParseBreathBlockEvents(header)
		if err != nil {
			log.Errorw("indexer: failed to parse breath block reward events", "err", err)
			return err
		}

		events = append(events, breatheEvents...)
	}

	err := i.store.SaveBreathBlockRewardEvents(ctx, events)
	if err != nil {
		log.Errorw("indexer: failed to save breath block reward events", "err", err)
		return err
	}

	return nil
}

// handleStakeCreditCalls
// eth_call for Validator, Delegator
func (i *indexer) handleStakeCreditCalls(ctx context.Context, header *types.Header) error {
	var validators []*model.Validator
	var delegators []*model.Delegator

	for _, v := range i.stakeCredits {
		validator, err := v.ParseValidator(header)
		if err != nil {
			log.Errorw("indexer: failed to parse validator reward", "err", err)
			return err
		}

		validators = append(validators, validator)

		for _, d := range i.delegators {
			delegator, err := v.ParseDelegator(header, d)
			if err != nil {
				log.Errorw("indexer: failed to parse delegator reward", "err", err)
				return err
			}

			delegators = append(delegators, delegator)
		}
	}

	err := syncutils.BatchRun(
		func() error {
			var err error
			err = i.store.SaveValidators(ctx, validators)
			if err != nil {
				log.Errorw("indexer: failed to save validators", "err", err)
				return err
			}

			return nil
		},

		func() error {
			var err error
			err = i.store.SaveDelegators(ctx, delegators)
			if err != nil {
				log.Errorw("indexer: failed to save delegators", "err", err)
				return err
			}

			return nil
		},
	)
	if err != nil {
		log.Errorw("indexer: failed to save stakeCredit calls data", "err", err)
		return err
	}

	return nil
}

func RtyDo(retryFn func() error) error {
	return retry.Do(
		retryFn,
		RtyAttempts,
		RtyDelay,
		RtyErr,
		retry.OnRetry(func(n uint, err error) {
			log.Errorf("indexer: retry failed, err: %v, attempt: %d times, max_attempts: %d", err, n+1, RtyNum)
		}),
	)
}

var BreatheBlockInterval uint64 = 86400 // Controls the interval for updateValidatorSetV2

// the params should be two blocks' time(timestamp)
func sameDayInUTC(first, second uint64) bool {
	return first/BreatheBlockInterval == second/BreatheBlockInterval
}

func isBreatheBlock(lastBlockTime, blockTime uint64) bool {
	return lastBlockTime != 0 && !sameDayInUTC(lastBlockTime, blockTime)
}
