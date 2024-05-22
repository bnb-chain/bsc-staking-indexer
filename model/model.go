package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ValidatorInfo is a model for the validator_info table.
// getLogs{toAddress:stakeHub, topic:ValidatorCreated}.
type ValidatorInfo struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Operator  common.Address `gorm:"column:operator;type:BINARY(20);uniqueIndex:idx_operator"`
	Consensus common.Address `gorm:"column:consensus;type:BINARY(20)"`
	Credit    common.Address `gorm:"column:credit;type:BINARY(20)"`
}

func (*ValidatorInfo) TableName() string {
	return "validator_info"
}

// Validator is a model for the validator table.
type Validator struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Operator common.Address `gorm:"column:operator;type:BINARY(20);uniqueIndex:idx_operator_date"`
	// RewardAfterCommission query StakeCredit rewardRecord at the end of the day
	RewardAfterCommission *Big `gorm:"column:reward_after_commission;type:VARBINARY(32)"`
	// TotalPooledBNB query StakeCredit totalPooledBNBRecord at the end of the day
	TotalPooledBNB *Big `gorm:"column:total_pooled_bnb;type:VARBINARY(32)"`
	// TotalCreditToken query StakeCredit _totalSupply at the previous block of the breathing block
	TotalCreditToken *Big `gorm:"column:total_credit_token;type:VARBINARY(32)"`

	Date int64 `gorm:"column:date;uniqueIndex:idx_operator_date"`
}

func (*Validator) TableName() string {
	return "validator"
}

// Delegator is a model for the delegator table.
// only record delegators from config.
type Delegator struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Delegator common.Address `gorm:"column:delegator;type:BINARY(20);uniqueIndex:idx_delegator_operator_date"`
	Operator  common.Address `gorm:"column:operator;type:BINARY(20);uniqueIndex:idx_delegator_operator_date"`
	// Amount query stakeCredit getPooledBNB interface at the previous block of the breathing block
	Amount *Big `gorm:"column:amount;type:VARBINARY(32)"`

	Date int64 `gorm:"column:date;index:idx_date;uniqueIndex:idx_delegator_operator_date"`
}

func (*Delegator) TableName() string {
	return "delegator"
}

type DelegateAction string

const (
	Delegate   DelegateAction = "delegate"
	UnDelegate DelegateAction = "undelegate"
	ReDelegate DelegateAction = "redelegate"
)

// DelegateTx is a model for the delegate_tx table.
// getLogs{toAddress:stakeHub, topic:Delegated/Undelegated/Redelegated}.
type DelegateTx struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Delegator   common.Address `gorm:"column:delegator;type:BINARY(20)"`
	SrcOperator common.Address `gorm:"column:src_operator;type:BINARY(20)"`
	DstOperator common.Address `gorm:"column:dst_operator;type:BINARY(20)"`
	Action      DelegateAction `gorm:"column:action"`
	Amount      *Big           `gorm:"column:amount;type:VARBINARY(32)"`
	TxHash      common.Hash    `gorm:"column:tx_hash;type:BINARY(32);uniqueIndex:idx_tx_hash"`
	Timestamp   int64          `gorm:"column:timestamp"`
}

func (*DelegateTx) TableName() string {
	return "delegate_tx"
}

// BreathBlockRewardEvent
// getLogs{toAddress:[credits], topics:RewardReceived}
type BreathBlockRewardEvent struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Credit                common.Address `gorm:"column:operator;type:BINARY(20);uniqueIndex:idx_credit_date"`
	RewardAfterCommission *Big           `gorm:"column:reward_after_commission;type:VARBINARY(32)"`
	Commission            *Big           `gorm:"column:commission;type:VARBINARY(32)"`

	Date int64 `gorm:"column:date;uniqueIndex:idx_credit_date"`
}

func (*BreathBlockRewardEvent) TableName() string {
	return "breath_block_reward_event"
}

// SlashEvent
// getLogs{toAddress:stakeHub, topics:ValidatorSlashed}
type SlashEvent struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Operator common.Address `gorm:"column:operator;type:BINARY(20);index:idx_operator"`
	Amount   *Big           `gorm:"column:amount;type:VARBINARY(32)"`
	TxHash   common.Hash    `gorm:"column:tx_hash;type:BINARY(32);uniqueIndex:idx_tx_hash"`

	Date int64 `gorm:"column:date;index:idx_date"`
}

func (*SlashEvent) TableName() string {
	return "slash_event"
}

// Cursor helps to record number cursor.
type Cursor struct {
	OneRowId bool  `gorm:"one_row_id;not null;default:true;primaryKey"`
	Number   int64 `gorm:"number"`
}

func (*Cursor) TableName() string {
	return "cursor"
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

func TruncateToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
