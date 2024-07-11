package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

// ValidatorInfo is a model for the validator_info table.
// getLogs{toAddress:stakeHub, topic:ValidatorCreated}.
type ValidatorInfo struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Operator  string `gorm:"column:operator;type:CHAR(42);uniqueIndex:idx_operator"`
	Consensus string `gorm:"column:consensus;type:CHAR(42)"`
	Credit    string `gorm:"column:credit;type:CHAR(42)"`
	Date      int64  `gorm:"column:date"`
	Number    int64  `gorm:"column:number;"`
}

func (*ValidatorInfo) TableName() string {
	return "validator_info"
}

// Validator is a model for the validator table.
type Validator struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Operator string `gorm:"column:operator;type:CHAR(42);uniqueIndex:idx_operator_date"`
	// RewardAfterCommission query StakeCredit rewardRecord at the end of the day
	RewardAfterCommission decimal.Decimal `gorm:"column:reward_after_commission;type:decimal(65,8)"`
	// TotalPooledBNB query StakeCredit totalPooledBNBRecord at the end of the day
	TotalPooledBNB decimal.Decimal `gorm:"column:total_pooled_bnb;type:decimal(65,8)"`
	// TotalCreditToken query StakeCredit _totalSupply at the previous block of the breathing block
	TotalCreditToken decimal.Decimal `gorm:"column:total_credit_token;type:decimal(65,8)"`

	Date int64 `gorm:"column:date;uniqueIndex:idx_operator_date"`
}

func (*Validator) TableName() string {
	return "validator"
}

// Delegator is a model for the delegator table.
// only record delegators from config.
type Delegator struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Delegator string `gorm:"column:delegator;type:CHAR(42);uniqueIndex:idx_delegator_operator_date"`
	Operator  string `gorm:"column:operator;type:CHAR(42);uniqueIndex:idx_delegator_operator_date"`
	// Amount query stakeCredit getPooledBNB interface at the previous block of the breathing block
	Amount decimal.Decimal `gorm:"column:amount;type:decimal(65,8)"`
	Date   int64           `gorm:"column:date;index:idx_date;uniqueIndex:idx_delegator_operator_date"`
	Number int64           `gorm:"column:number;"`
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

	Delegator   string          `gorm:"column:delegator;type:CHAR(42)"`
	SrcOperator string          `gorm:"column:src_operator;type:CHAR(42)"`
	DstOperator string          `gorm:"column:dst_operator;type:CHAR(42)"`
	Action      DelegateAction  `gorm:"column:action;type:VARCHAR(16)"`
	Amount      decimal.Decimal `gorm:"column:amount;type:decimal(65,8)"`
	TxHash      string          `gorm:"column:tx_hash;type:CHAR(66);uniqueIndex:idx_tx_hash"`
	Timestamp   int64           `gorm:"column:timestamp"`
}

func (*DelegateTx) TableName() string {
	return "delegate_tx"
}

// BreathBlockRewardEvent
// getLogs{toAddress:[credits], topics:RewardReceived}
type BreathBlockRewardEvent struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Operator              string          `gorm:"column:operator;type:CHAR(42);uniqueIndex:idx_operator_date"`
	RewardAfterCommission decimal.Decimal `gorm:"column:reward_after_commission;type:decimal(65,8)"`
	Commission            decimal.Decimal `gorm:"column:commission;type:decimal(65,8)"`
	Date                  int64           `gorm:"column:date;uniqueIndex:idx_operator_date"`
	Number                int64           `gorm:"column:number;"`
}

func (*BreathBlockRewardEvent) TableName() string {
	return "breath_block_reward_event"
}

// SlashEvent
// getLogs{toAddress:stakeHub, topics:ValidatorSlashed}
type SlashEvent struct {
	ID uint64 `gorm:"column:id;primaryKey"`

	Operator string          `gorm:"column:operator;type:CHAR(42);index:idx_operator"`
	Amount   decimal.Decimal `gorm:"column:amount;type:decimal(65,8)"`
	TxHash   string          `gorm:"column:tx_hash;type:CHAR(66);uniqueIndex:idx_tx_hash"`
	Date     int64           `gorm:"column:date;index:idx_date"`
	Number   int64           `gorm:"column:number;"`
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
