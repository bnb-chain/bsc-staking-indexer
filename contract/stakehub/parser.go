package stakehub

import (
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bnb-chain/bsc-staking-indexer/model"
	"github.com/bnb-chain/bsc-staking-indexer/util/log"
	"github.com/bnb-chain/bsc-staking-indexer/util/syncutil"
)

var Address = common.HexToAddress("0x0000000000000000000000000000000000002002")

func New(client *ethclient.Client) (*Contract, error) {
	return NewContract(Address, client)
}

func (_contract *Contract) ParseValidatorInfos(number uint64) ([]*model.ValidatorInfo, error) {
	iterator, err := _contract.FilterValidatorCreated(&bind.FilterOpts{Start: number, End: &number},
		nil, nil, nil)

	if err != nil {
		log.Errorw("failed to filter validator created", "err", err)
		return nil, err
	}

	validatorInfos := make([]*model.ValidatorInfo, 0)
	defer iterator.Close()

	for iterator.Next() {
		event := iterator.Event

		validatorInfos = append(validatorInfos, &model.ValidatorInfo{
			Operator:  event.OperatorAddress,
			Consensus: event.ConsensusAddress,
			Credit:    event.CreditContract,
		})
	}

	return validatorInfos, nil
}

func (_contract *Contract) ParseDelegateTxs(header *types.Header, delegator []common.Address) ([]*model.DelegateTx, error) {
	number := header.Number.Uint64()
	var (
		delegateIterator   *ContractDelegatedIterator
		unDelegateIterator *ContractUndelegatedIterator
		reDelegateIterator *ContractRedelegatedIterator
	)

	err := syncutil.BatchRun(
		func() error {
			var err error
			delegateIterator, err = _contract.FilterDelegated(&bind.FilterOpts{Start: number, End: &number},
				nil, delegator)

			if err != nil {
				log.Errorw("failed to filter delegated", "err", err)
				return err
			}

			return nil
		},
		func() error {
			var err error
			unDelegateIterator, err = _contract.FilterUndelegated(&bind.FilterOpts{Start: number, End: &number},
				nil, delegator)

			if err != nil {
				log.Errorw("failed to filter undelegated", "err", err)
				return err
			}

			return nil
		},
		func() error {
			var err error
			reDelegateIterator, err = _contract.FilterRedelegated(&bind.FilterOpts{Start: number, End: &number},
				nil, nil, delegator)

			if err != nil {
				log.Errorw("failed to filter redelegated", "err", err)
				return err
			}

			return nil
		},
	)
	if err != nil {
		log.Errorw("failed to filter delegate txs", "err", err)
		return nil, err
	}

	delegateTxs := make([]*model.DelegateTx, 0)

	defer delegateIterator.Close()
	for delegateIterator.Next() {
		event := delegateIterator.Event

		delegateTxs = append(delegateTxs, &model.DelegateTx{
			Delegator:   event.Delegator,
			SrcOperator: event.OperatorAddress,
			Action:      model.Delegate,
			Amount:      (*model.Big)(event.BnbAmount),
			TxHash:      event.Raw.TxHash,
			Timestamp:   int64(header.Time),
		})
	}

	defer unDelegateIterator.Close()
	for unDelegateIterator.Next() {
		event := unDelegateIterator.Event

		delegateTxs = append(delegateTxs, &model.DelegateTx{
			Delegator:   event.Delegator,
			SrcOperator: event.OperatorAddress,
			Action:      model.UnDelegate,
			Amount:      (*model.Big)(event.BnbAmount),
			TxHash:      event.Raw.TxHash,
			Timestamp:   int64(header.Time),
		})
	}

	defer reDelegateIterator.Close()
	for reDelegateIterator.Next() {
		event := reDelegateIterator.Event

		delegateTxs = append(delegateTxs, &model.DelegateTx{
			Delegator:   event.Delegator,
			SrcOperator: event.SrcValidator,
			DstOperator: event.DstValidator,
			Action:      model.ReDelegate,
			Amount:      (*model.Big)(event.BnbAmount),
			TxHash:      event.Raw.TxHash,
			Timestamp:   int64(header.Time),
		})
	}

	return delegateTxs, nil
}

func (_contract *Contract) ParseSlashEvents(header *types.Header) ([]*model.SlashEvent, error) {
	number := header.Number.Uint64()
	iterator, err := _contract.FilterValidatorSlashed(&bind.FilterOpts{Start: number, End: &number},
		nil)

	if err != nil {
		log.Errorw("failed to filter validator slashed", "err", err)
		return nil, err
	}

	slashEvents := make([]*model.SlashEvent, 0)
	defer iterator.Close()

	for iterator.Next() {
		event := iterator.Event

		slashEvents = append(slashEvents, &model.SlashEvent{
			Operator: event.OperatorAddress,
			Amount:   (*model.Big)(event.SlashAmount),
			TxHash:   event.Raw.TxHash,
			Date:     model.TruncateToDate(time.Unix(int64(header.Time), 0)).Unix(),
		})
	}

	return slashEvents, nil
}
