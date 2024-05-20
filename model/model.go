package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ValidatorInfo is a model for the validator_info table.
// query stakeHub getValidators at startup only once and only their validator?
type ValidatorInfo struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Operator  common.Address `gorm:"column:operator;uniqueIndex:idx_operator"`
	Consensus common.Address `gorm:"column:consensus"`
	Credit    common.Address `gorm:"column:credit"`
}

// Validator is a model for the validator table.
type Validator struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Operator common.Address `gorm:"column:operator;uniqueIndex:idx_operator_date"`
	// RewardAfterCommission query StakeCredit rewardRecord at the end of the day
	RewardAfterCommission *Big `gorm:"column:reward_after_commission;type:VARBINARY(32)"`
	// TotalPooledBNB query StakeCredit totalPooledBNBRecord at the end of the day
	TotalPooledBNB *Big `gorm:"column:total_pooled_bnb;type:VARBINARY(32)"`
	// TotalCreditToken query StakeCredit _totalSupply at the previous block of the breathing block
	TotalCreditToken *Big `gorm:"column:total_credit_token;type:VARBINARY(32)"`

	Date int64 `gorm:"column:date;uniqueIndex:idx_operator_date"`
}

// Delegator is a model for the delegator table.
// only record their delegator from config.
// need listen stakeHub Delegated event?
type Delegator struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Delegator common.Address `gorm:"column:delegator"`
	Operator  common.Address `gorm:"column:operator"`
	// Amount query stakeCredit getPooledBNB interface at the previous block of the breathing block
	Amount *Big `gorm:"column:amount;type:VARBINARY(32)"`

	Date int64 `gorm:"column:date;index:idx_date"`
}

type DelegateAction string

const (
	Delegate    DelegateAction = "delegate"
	UnDelegate  DelegateAction = "undelegate"
	ReDelegator DelegateAction = "redelegate"
)

// DelegateTx is a model for the delegate_tx table.
// getLogs{toAddress:stakeHub, topic:Delegated/Undelegated/Redelegated}.
type DelegateTx struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Delegator   common.Address `gorm:"column:delegator"`
	SrcOperator common.Address `gorm:"column:src_operator"`
	DstOperator common.Address `gorm:"column:dst_operator"`
	Action      DelegateAction `gorm:"column:action"`
	Amount      *Big           `gorm:"column:amount;type:VARBINARY(32)"`
	TxHash      common.Hash    `gorm:"column:tx_hash;uniqueIndex:idx_tx_hash"`
	Timestamp   int64          `gorm:"column:timestamp"`
}

// BreathBlockRewardEvent
// getLogs{toAddress:[credits], topics:RewardReceived}
type BreathBlockRewardEvent struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Credit                common.Address `gorm:"column:operator;uniqueIndex:idx_credit_date"`
	RewardAfterCommission *Big           `gorm:"column:reward_after_commission;type:VARBINARY(32)"`
	Commission            *Big           `gorm:"column:commission;type:VARBINARY(32)"`

	Date int64 `gorm:"column:date;uniqueIndex:idx_credit_date"`
}

// SlashEvent
// getLogs{toAddress:[stakeHub], topics:ValidatorSlashed}
type SlashEvent struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Operator common.Address `gorm:"column:operator;index:idx_operator"`
	Amount   *Big           `gorm:"column:amount;type:VARBINARY(32)"`
	TxHash   common.Hash    `gorm:"column:tx_hash;uniqueIndex:idx_tx_hash"`

	Date int64 `gorm:"column:date;index:idx_date"`
}

// Cursor helps to record number cursor.
type Cursor struct {
	OneRowId bool  `gorm:"one_row_id;not null;default:true;primaryKey"`
	Number   int64 `gorm:"number"`
}

type Big big.Int

func (i *Big) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal Big value:", value))
	}

	i.Raw().SetBytes(bytes)
	return nil
}

func (i *Big) Value() (driver.Value, error) {
	return i.Raw().Bytes(), nil
}

func (i *Big) Raw() *big.Int {
	return (*big.Int)(i)
}

func (i *Big) MarshalText() ([]byte, error) {
	return []byte(hexutil.EncodeBig((*big.Int)(i))), nil
}
