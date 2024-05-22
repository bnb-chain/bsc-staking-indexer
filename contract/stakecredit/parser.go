package stakecredit

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bnb-chain/bsc-staking-indexer/model"
	"github.com/bnb-chain/bsc-staking-indexer/util/log"
)

const BreatheBlockInterval = 86400

type ContractWithInfo struct {
	*Contract
	credit   common.Address
	operator common.Address
}

func New(operator, credit common.Address, client *ethclient.Client) (*ContractWithInfo, error) {
	contract, err := NewContract(operator, client)
	if err != nil {
		return nil, err
	}

	return &ContractWithInfo{Contract: contract, operator: operator, credit: credit}, nil
}

func (_contract *ContractWithInfo) ParseBreathBlockEvents(header *types.Header) ([]*model.BreathBlockRewardEvent, error) {
	number := header.Number.Uint64()
	iterator, err := _contract.FilterRewardReceived(&bind.FilterOpts{Start: number, End: &number})
	if err != nil {
		log.Errorw("failed to filter reward received", "err", err)
		return nil, err
	}

	events := make([]*model.BreathBlockRewardEvent, 0)
	defer iterator.Close()

	for iterator.Next() {
		event := iterator.Event
		events = append(events, &model.BreathBlockRewardEvent{
			Credit:                _contract.credit,
			RewardAfterCommission: (*model.Big)(event.RewardToAll),
			Commission:            (*model.Big)(event.Commission),
			Date:                  model.TruncateToDate(time.Unix(int64(header.Time), 0)).Unix(),
		})
	}

	return events, nil
}

func (_contract *ContractWithInfo) ParseValidator(header *types.Header) (*model.Validator, error) {
	reward, err := _contract.RewardRecord(&bind.CallOpts{BlockNumber: header.Number},
		big.NewInt(int64(header.Time/BreatheBlockInterval)-1))
	if err != nil {
		log.Errorw("failed to call reward record", "err", err)
		return nil, err
	}

	bnb, err := _contract.TotalPooledBNBRecord(&bind.CallOpts{BlockNumber: header.Number},
		big.NewInt(int64(header.Time/BreatheBlockInterval)-1))
	if err != nil {
		log.Errorw("failed to call total pooled bnb record", "err", err)
		return nil, err
	}

	credit, err := _contract.TotalSupply(&bind.CallOpts{BlockNumber: big.NewInt(header.Number.Int64() - 1)})
	if err != nil {
		log.Errorw("failed to call total supply", "err", err)
		return nil, err
	}

	return &model.Validator{
		Operator:              _contract.operator,
		RewardAfterCommission: (*model.Big)(reward),
		TotalPooledBNB:        (*model.Big)(bnb),
		TotalCreditToken:      (*model.Big)(credit),

		Date: model.TruncateToDate(time.Unix(int64(header.Time), 0).AddDate(0, 0, -1)).Unix(),
	}, nil
}

func (_contract *ContractWithInfo) ParseDelegator(header *types.Header, delegator common.Address) (*model.Delegator, error) {
	amount, err := _contract.GetPooledBNB(&bind.CallOpts{BlockNumber: big.NewInt(header.Number.Int64() - 1)}, delegator)
	if err != nil {
		log.Errorw("failed to call get pooled bnb", "err", err)
		return nil, err
	}

	return &model.Delegator{
		Delegator: delegator,
		Operator:  _contract.operator,
		Amount:    (*model.Big)(amount),

		Date: model.TruncateToDate(time.Unix(int64(header.Time), 0).AddDate(0, 0, -1)).Unix(),
	}, nil
}
