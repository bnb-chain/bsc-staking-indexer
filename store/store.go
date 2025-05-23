package store

import (
	"context"

	"github.com/node-real/go-pkg/mysqlclient"
	"github.com/node-real/go-pkg/utils/syncutils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/bnb-chain/bsc-staking-indexer/model"
)

type Store interface {
	SaveValidatorInfos(ctx context.Context, infos []*model.ValidatorInfo) error
	SaveDelegateTxs(ctx context.Context, txs []*model.DelegateTx) error
	SaveSlashEvents(ctx context.Context, events []*model.SlashEvent) error
	SaveBreathBlockRewardEvents(ctx context.Context, events []*model.BreathBlockRewardEvent) error
	SaveValidators(ctx context.Context, validators []*model.Validator) error
	SaveDelegators(ctx context.Context, delegators []*model.Delegator) error
	SaveCursor(ctx context.Context, number int64) error

	QueryValidatorInfos(ctx context.Context) ([]*model.ValidatorInfo, error)
	QueryCursor(ctx context.Context) (int64, error)

	QueryBreathBlockRewardEvent(ctx context.Context, operator string, fromDate, toDate int64) ([]*model.BreathBlockRewardEvent, error)
	QueryDelegator(ctx context.Context, delegator, operator string, startDateOfNextMonth int64) (*model.Delegator, error)
}

type store struct {
	db *gorm.DB
}

func NewStore(cfg mysqlclient.Config) (Store, error) {
	db, err := mysqlclient.New(&cfg)
	if err != nil {
		return nil, err
	}

	err = syncutils.BatchRun(
		func() error {
			return db.Table((&model.ValidatorInfo{}).TableName()).AutoMigrate(&model.ValidatorInfo{})
		},
		func() error { return db.Table((&model.DelegateTx{}).TableName()).AutoMigrate(&model.DelegateTx{}) },
		func() error { return db.Table((&model.SlashEvent{}).TableName()).AutoMigrate(&model.SlashEvent{}) },
		func() error {
			return db.Table((&model.BreathBlockRewardEvent{}).TableName()).AutoMigrate(&model.BreathBlockRewardEvent{})
		},
		func() error { return db.Table((&model.Validator{}).TableName()).AutoMigrate(&model.Validator{}) },
		func() error { return db.Table((&model.Delegator{}).TableName()).AutoMigrate(&model.Delegator{}) },
		func() error { return db.Table((&model.Cursor{}).TableName()).AutoMigrate(&model.Cursor{}) },
	)
	if err != nil {
		return nil, err
	}

	return &store{
		db: db,
	}, nil
}

func (s *store) SaveValidatorInfos(ctx context.Context, infos []*model.ValidatorInfo) error {
	if len(infos) == 0 {
		return nil
	}

	return s.db.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(infos).Error
}

func (s *store) SaveDelegateTxs(ctx context.Context, txs []*model.DelegateTx) error {
	if len(txs) == 0 {
		return nil
	}

	return s.db.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(txs).Error
}

func (s *store) SaveSlashEvents(ctx context.Context, events []*model.SlashEvent) error {
	if len(events) == 0 {
		return nil
	}

	return s.db.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(events).Error
}

func (s *store) SaveBreathBlockRewardEvents(ctx context.Context, events []*model.BreathBlockRewardEvent) error {
	if len(events) == 0 {
		return nil
	}

	return s.db.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(events).Error
}

func (s *store) SaveValidators(ctx context.Context, validators []*model.Validator) error {
	if len(validators) == 0 {
		return nil
	}

	return s.db.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(validators).Error
}

func (s *store) SaveDelegators(ctx context.Context, delegators []*model.Delegator) error {
	if len(delegators) == 0 {
		return nil
	}

	return s.db.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(delegators).Error
}

func (s *store) SaveCursor(ctx context.Context, number int64) error {
	cursor := model.Cursor{Number: number}
	return s.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "one_row_id"}}, UpdateAll: true}).Save(&cursor).Error
}

func (s *store) QueryValidatorInfos(ctx context.Context) ([]*model.ValidatorInfo, error) {
	var infos []*model.ValidatorInfo
	err := s.db.WithContext(ctx).Model(&model.ValidatorInfo{}).Scan(&infos).Error

	return infos, err
}

func (s *store) QueryCursor(ctx context.Context) (int64, error) {
	var cursor model.Cursor
	err := s.db.WithContext(ctx).Model(&model.Cursor{}).First(&cursor).Error

	if err == gorm.ErrRecordNotFound {
		err = nil
	}

	return cursor.Number, err
}

func (s *store) QueryBreathBlockRewardEvent(ctx context.Context, operator string, fromDate, toDate int64) (
	[]*model.BreathBlockRewardEvent, error) {
	var events []*model.BreathBlockRewardEvent
	err := s.db.WithContext(ctx).Model(&model.BreathBlockRewardEvent{}).Where(
		"operator = ? AND date >= ? AND date < ?", operator, fromDate, toDate).Scan(&events).Error
	return events, err
}

func (s *store) QueryDelegator(ctx context.Context, delegator, operator string, startDateOfNextMonth int64) (
	*model.Delegator, error) {
	var res *model.Delegator
	err := s.db.WithContext(ctx).Model(&model.Delegator{}).Where(
		"delegator = ? AND operator = ? AND date < ?",
		delegator, operator, startDateOfNextMonth).Order("date desc").Limit(1).Find(&res).Error
	return res, err
}
