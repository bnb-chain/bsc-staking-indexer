package indexer

import (
	"context"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bnb-chain/bsc-staking-indexer/log"
)

var (
	RtyNum      = uint(5)
	RtyAttempts = retry.Attempts(RtyNum)
	RtyDelay    = retry.Delay(time.Millisecond * 50)
	RtyErr      = retry.LastErrorOnly(true)
)

type Indexer interface {
	Start(ctx context.Context)
}

type indexer struct {
	cursor                    int64
	NumberOfBlocksForFinality int64

	client *ethclient.Client
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
	err := retry.Do(
		func() error {
			var err error
			finalizedHeader, err = i.client.FinalizedHeader(ctx, i.NumberOfBlocksForFinality)
			return err
		},
		RtyAttempts,
		RtyDelay,
		RtyErr,
		retry.OnRetry(func(n uint, err error) {
			log.Errorf("indexer: failed to get finalized header, attempt: %d times, max_attempts: %d", n+1, RtyNum)
		}),
	)
	if err != nil {
		log.Errorw("indexer: failed to get finalized header", "err", err)
		return err
	}

	start := i.cursor + 1
	end := finalizedHeader.Number.Int64()

	for number := start; number <= end; number++ {
		// getLogs for ValidatorInfo, DelegatorTx, SlashEvent

		// if breath block, getLogs for BreathBlockRewardEvent

		// if breath block, query for Validator, Delegator
	}

	return nil
}
