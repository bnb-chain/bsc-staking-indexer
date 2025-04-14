package indexer

import (
	"context"
	"crypto/tls"
	"math/big"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/bnb-chain/bsc-mev-sentry/syncutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/node-real/go-pkg/log"

	"github.com/bnb-chain/bsc-staking-indexer/contract/stakecredit"
	"github.com/bnb-chain/bsc-staking-indexer/contract/stakehub"
	"github.com/bnb-chain/bsc-staking-indexer/model"
	"github.com/bnb-chain/bsc-staking-indexer/store"
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

	concurrentRPCCount int64 = 200
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

	delegators := make(map[common.Address]bool)
	stakeCredits := make(map[common.Address]*stakecredit.ContractWithInfo)
	stakeCreditAddresses := make([]common.Address, 0)
	for _, v := range validatorInfos {
		stakeCredits[common.HexToAddress(v.Credit)], err = stakecredit.New(
			common.HexToAddress(v.Operator), common.HexToAddress(v.Credit), client)
		if err != nil {
			log.Errorw("indexer: failed to new stakecredit", "err", err)
			return nil, err
		}
		stakeCreditAddresses = append(stakeCreditAddresses, common.HexToAddress(v.Credit))
		delegators[common.HexToAddress(v.Operator)] = true
	}

	for _, v := range cfg.Delegators {
		delegators[v] = true
	}

	return &indexer{
		cursor:                    cursor,
		NumberOfBlocksForFinality: cfg.NumberOfBlocksForFinality,
		delegators:                delegators,
		store:                     store,
		client:                    client,
		stakeHub:                  stakeHub,
		stakeCredits:              stakeCredits,
		stakeCreditAddresses:      stakeCreditAddresses,
		headers:                   make(chan chan *headerLoadResult, concurrentRPCCount),
	}, nil
}

type headerLoadResult struct {
	header                  *types.Header
	validatorInfos          []*model.ValidatorInfo
	delegateTxs             []*model.DelegateTx
	slashEvents             []*model.SlashEvent
	breathBlockRewardEvents []*model.BreathBlockRewardEvent
	validators              []*model.Validator
	delegators              []*model.Delegator

	err error
}

type indexer struct {
	cursor                    int64
	NumberOfBlocksForFinality int64
	mu                        sync.RWMutex
	delegators                map[common.Address]bool
	parentHeader              *types.Header

	store        store.Store
	client       *ethclient.Client
	stakeHub     *stakehub.Contract
	stakeCredits map[common.Address]*stakecredit.ContractWithInfo // credit address -> stake credit

	stakeCreditAddresses []common.Address

	headers chan chan *headerLoadResult
}

func (i *indexer) Start(ctx context.Context) {
	ticker := time.NewTicker(1500 * time.Millisecond)
	defer ticker.Stop()

	go i.index(ctx)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := i.fetchBlockData(ctx)
			if err != nil {
				log.Errorw("indexer: failed to fetch headers", "err", err)
			}
		}
	}
}

func (i *indexer) index(ctx context.Context) {
	for {
		select {
		case headerCh, ok := <-i.headers:
			if !ok {
				log.Info("indexer: headers chan closed")
			}

			headerRes := <-headerCh
			close(headerCh)

			if headerRes.err != nil {
				log.Panicw("indexer: failed to get header", "err", headerRes.err)
				return
			}

			header := headerRes.header
			if header == nil {
				log.Panicw("indexer: header is nil")
				return
			}

			if i.parentHeader == nil {
				if err := RtyDo(func() error {
					var er error
					i.parentHeader, er = i.client.HeaderByNumber(ctx, big.NewInt(header.Number.Int64()-1))
					return er
				}); err != nil {
					log.Panicw("indexer: failed to get header by number", "number", header.Number.Int64()-1, "err", err)
					return
				}
			}

			err := i.fetchStakeCreditLogsData(headerRes)
			if err != nil {
				log.Errorw("indexer: failed to fetch stakeCredit data", "err", err)
				headerRes.err = err
				headerCh <- headerRes
				return
			}

			if isBreatheBlock(i.parentHeader.Time, header.Time) {
				err = i.fetchStakeCreditCallData(headerRes)
				if err != nil {
					log.Panicw("indexer: failed to process validator and delegator", "err", err)
					return
				}
			}

			err = i.saveBlockData(ctx, headerRes)
			for err != nil {
				err = i.saveBlockData(ctx, headerRes)
			}

			log.Infow("indexer: processed block", "number", header.Number.Int64(), "hash", header.Hash().Hex())
			i.parentHeader = header

		case <-time.After(500 * time.Millisecond):
			log.Info("indexer: no headers to process")

		case <-ctx.Done():
			log.Info("indexer: context done")
			return
		}
	}
}

func (i *indexer) saveBlockData(ctx context.Context, headerRes *headerLoadResult) (err error) {
	defer recordLatency("saveBlockData", time.Now())

	err = syncutils.BatchRun(
		func() error { return i.store.SaveValidatorInfos(ctx, headerRes.validatorInfos) },
		func() error { return i.store.SaveDelegateTxs(ctx, headerRes.delegateTxs) },
		func() error { return i.store.SaveSlashEvents(ctx, headerRes.slashEvents) },
		func() error { return i.store.SaveBreathBlockRewardEvents(ctx, headerRes.breathBlockRewardEvents) },
		func() error { return i.store.SaveValidators(ctx, headerRes.validators) },
		func() error { return i.store.SaveDelegators(ctx, headerRes.delegators) },
	)

	if err != nil {
		log.Errorw("indexer: failed to save block logs data", "number", headerRes.header.Number.Int64(), "err", err)
		return err
	}

	err = i.store.SaveCursor(ctx, headerRes.header.Number.Int64())
	if err != nil {
		log.Errorw("indexer: failed to save cursor", "number", headerRes.header.Number.Int64(), "err", err)
		return err
	}

	return
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

func bloomFilter(bloom types.Bloom, addresses []common.Address, topics [][]common.Hash) bool {
	if len(addresses) > 0 {
		var included bool
		for _, addr := range addresses {
			if types.BloomLookup(bloom, addr) {
				included = true
				break
			}
		}
		if !included {
			return false
		}
	}

	for _, sub := range topics {
		included := len(sub) == 0 // empty rule set == wildcard
		for _, topic := range sub {
			if types.BloomLookup(bloom, topic) {
				included = true
				break
			}
		}
		if !included {
			return false
		}
	}
	return true
}

func recordLatency(method string, start time.Time) {
	log.Infow("time cost", "method", method, "time", time.Since(start).Milliseconds())
}
