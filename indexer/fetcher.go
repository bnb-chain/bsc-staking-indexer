package indexer

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/node-real/go-pkg/log"
	"github.com/node-real/go-pkg/utils/syncutils"

	"github.com/bnb-chain/bsc-staking-indexer/contract/stakecredit"
	"github.com/bnb-chain/bsc-staking-indexer/contract/stakehub"
	"github.com/bnb-chain/bsc-staking-indexer/model"
)

func (i *indexer) fetchBlockData(ctx context.Context) error {
	var finalizedHeader *types.Header

	if err := RtyDo(func() error {
		var er error
		finalizedHeader, er = i.client.FinalizedHeader(ctx, i.NumberOfBlocksForFinality)
		return er
	}); err != nil {
		log.Errorw("indexer: failed to get finalized header", "err", err)
		return err
	}

	start := i.cursor + 1
	end := finalizedHeader.Number.Int64()
	if end > start+concurrentRPCCount-1 {
		end = start + concurrentRPCCount - 1
	}

	var wg sync.WaitGroup
	wg.Add(int(end - start + 1))

	for number := start; number <= end; number++ {
		headerCh := make(chan *headerLoadResult, 1)
		i.headers <- headerCh

		go func(number int64) {
			defer wg.Done()

			var header *types.Header

			err := RtyDo(func() error {
				var er error
				header, er = i.client.HeaderByNumber(ctx, big.NewInt(number))
				return er
			})

			headerRes := &headerLoadResult{header: header}

			if err != nil {
				log.Errorw("indexer: failed to get header by number", "number", number, "err", err)
				headerRes.err = err
				headerCh <- headerRes
				return
			}

			err = i.fetchStakeHubLogsData(headerRes)
			if err != nil {
				log.Errorw("indexer: failed to fetch stakeHub data", "err", err)
				headerRes.err = err
				headerCh <- headerRes
				return
			}

			headerCh <- headerRes
		}(number)
	}

	wg.Wait()
	i.cursor = end

	return nil
}

// fetchStakeHubLogsData
func (i *indexer) fetchStakeHubLogsData(headerRes *headerLoadResult) error {
	defer recordLatency("fetchStakeHubLogsData", time.Now())
	header := headerRes.header

	if !bloomFilter(header.Bloom, []common.Address{stakehub.Address}, [][]common.Hash{stakehub.Topics}) {
		return nil
	}

	var validatorInfos []*model.ValidatorInfo
	var slashEvents []*model.SlashEvent

	err := syncutils.BatchRun(
		func() error {
			err := RtyDo(func() error {
				var er error
				validatorInfos, er = i.stakeHub.ParseValidatorInfos(header)
				return er
			})
			if err != nil {
				log.Errorw("indexer: failed to parse validator created", "err", err)
				return err
			}
			return nil
		},
		func() error {
			err := RtyDo(func() error {
				var er error
				slashEvents, er = i.stakeHub.ParseSlashEvents(header)
				return er
			})
			if err != nil {
				log.Errorw("indexer: failed to parse slash events", "err", err)
				return err
			}
			return nil
		},
	)

	if err != nil {
		log.Errorw("indexer: failed to fetch stakeHub data", "err", err)
		return err
	}

	delegators := make([]common.Address, 0)

	i.mu.Lock()
	for _, v := range validatorInfos {
		i.delegators[common.HexToAddress(v.Operator)] = true
	}
	for k := range i.delegators {
		delegators = append(delegators, k)
	}
	i.mu.Unlock()

	delegateTxs, er := i.stakeHub.ParseDelegateTxs(header, delegators)
	if er != nil {
		log.Errorw("indexer: failed to parse delegated", "err", err)
		return err
	}

	headerRes.validatorInfos = validatorInfos
	headerRes.delegateTxs = delegateTxs
	headerRes.slashEvents = slashEvents

	return nil
}

// fetchStakeCreditLogsData
func (i *indexer) fetchStakeCreditLogsData(headerRes *headerLoadResult) (err error) {
	defer recordLatency("fetchStakeCreditLogsData", time.Now())
	header := headerRes.header

	for _, v := range headerRes.validatorInfos {
		if i.stakeCredits[common.HexToAddress(v.Credit)] == nil {
			credit, err := stakecredit.New(common.HexToAddress(v.Operator), common.HexToAddress(v.Credit), i.client)
			if err != nil {
				log.Errorw("indexer: failed to new stakecredit", "err", err)
			}
			i.stakeCredits[common.HexToAddress(v.Credit)] = credit
			i.stakeCreditAddresses = append(i.stakeCreditAddresses, common.HexToAddress(v.Credit))
		}
	}

	if !bloomFilter(header.Bloom, i.stakeCreditAddresses, [][]common.Hash{stakecredit.Topics}) {
		return nil
	}

	log.Infow("operator sum", "count", len(i.stakeCredits))

	queue := make(chan *model.BreathBlockRewardEvent, len(i.stakeCredits))
	batchRunner := syncutils.NewBatchRunner()
	for _, v := range i.stakeCredits {
		v := v
		batchRunner.AddTasks(func() error {
			var breatheEvents []*model.BreathBlockRewardEvent
			err = RtyDo(func() error {
				var er error
				breatheEvents, er = v.ParseBreathBlockEvents(header)
				return er
			})
			if err != nil {
				log.Errorw("indexer: failed to parse breath block reward events", "err", err)
				return err
			}

			for _, e := range breatheEvents {
				queue <- e
			}

			return nil
		})
	}
	err = batchRunner.Exec()
	if err != nil {
		log.Errorw("indexer: failed to fetch stakeCredit logs data", "err", err)
		return err
	}

	close(queue)

	breathBlockRewardEvents := make([]*model.BreathBlockRewardEvent, 0)
	for v := range queue {
		breathBlockRewardEvents = append(breathBlockRewardEvents, v)
	}

	headerRes.breathBlockRewardEvents = breathBlockRewardEvents
	log.Infow("breathBlockRewardEvents", "count", len(breathBlockRewardEvents))

	return nil
}

// fetchStakeCreditCallData
// eth_call for Validator, Delegator
func (i *indexer) fetchStakeCreditCallData(headerRes *headerLoadResult) (err error) {
	defer recordLatency("fetchStakeCreditCallData", time.Now())
	header := headerRes.header

	var validators []*model.Validator
	var delegators []*model.Delegator

	delegator := make([]common.Address, 0)
	i.mu.Lock()
	for k := range i.delegators {
		delegator = append(delegator, k)
	}
	i.mu.Unlock()

	for _, v := range i.stakeCredits {
		var validator *model.Validator
		err = RtyDo(func() error {
			var er error
			validator, er = v.ParseValidator(header)
			return er
		})
		if err != nil {
			log.Errorw("indexer: failed to parse validator reward", "err", err)
			return err
		}

		validators = append(validators, validator)

		for _, d := range delegator {
			var delegator *model.Delegator
			err = RtyDo(func() error {
				var er error
				delegator, er = v.ParseDelegator(header, d)
				return er
			})
			if err != nil {
				log.Errorw("indexer: failed to parse delegator reward", "err", err)
				return err
			}

			delegators = append(delegators, delegator)
		}
	}

	headerRes.validators = validators
	headerRes.delegators = delegators

	return nil
}
