package mysqlclient

// copy from github.com/node-real/blocktree/pkg/mysqlclient

import (
	"context"
	"database/sql"
	"errors"
	"time"

	mysqlDriver "github.com/node-real/go-mysql-driver"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/node-real/go-pkg/aws/secrets"
	"github.com/node-real/go-pkg/log"
	"github.com/node-real/go-pkg/units"
)

type Config struct {
	DSN                string
	Secrets            *secrets.Params
	SlowThreshold      units.Duration
	MaxIdleConns       int
	MaxOpenConns       int
	ConnMaxIdleTime    units.Duration
	ConnMaxLifetime    units.Duration
	EnableAWSConnector bool // use aws-rds-mysql driver
	LocalAZFirst       bool // use db instance in local az
	UseWriter          bool // use db writer instance
	DisableReader      bool // do not use db reader instance, default is false
	FallbackToWriter   bool // fallback to use writer instance, if no reader is found
	DisableMetrics     bool // disable metrics
	DisableTracing     bool
}

func New(cfg *Config) (*gorm.DB, error) {
	if cfg.Secrets != nil {
		secret, err := secrets.GetString(cfg.Secrets)
		if err != nil {
			log.Errorf("invalid secrets %+v err:%v", cfg.Secrets, err)
			return nil, err
		}
		cfg.DSN = secret
	}

	mysqlConfig := mysql.Config{
		DSN: cfg.DSN,
	}

	if cfg.EnableAWSConnector {
		mysqlConfig.DriverName = "aws-rds-mysql"
		driverConfig, err := mysqlDriver.ParseDSN(cfg.DSN)
		if err != nil {
			return nil, err
		}
		if cfg.LocalAZFirst {
			driverConfig.LocalAZFirst = true
		}
		if cfg.UseWriter {
			driverConfig.UseWriter = true
		}
		if cfg.DisableReader {
			driverConfig.DisableReader = true
		}
		if cfg.FallbackToWriter {
			driverConfig.FallbackToWriter = true
		}
		if cfg.DisableMetrics {
			driverConfig.DisableMetrics = true
		}
		mysqlConfig.DSN = driverConfig.FormatDSN()
	} else {
		dsn, err := mysqlDriver.FormatDSNSafe(cfg.DSN)
		if err != nil {
			return nil, err
		}
		mysqlConfig.DSN = dsn
	}

	db, err := gorm.Open(mysql.New(mysqlConfig),
		&gorm.Config{
			Logger:                                   &loggerAdaptor{slowThreshold: time.Duration(cfg.SlowThreshold)},
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
	if err != nil {
		log.Errorw("failed to open database", "err", err)
		return nil, err
	}

	if !cfg.DisableTracing {
		if err := db.Use(otelgorm.NewPlugin()); err != nil {
			log.Errorw("failed to register gorm tracing plugin", "err", err)
			return nil, err
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Errorw("failed to get database", "err", err)
		return nil, err
	}

	if cfg.MaxOpenConns <= 0 {
		cfg.MaxOpenConns = 256
	}
	if cfg.MaxIdleConns <= 0 {
		cfg.MaxIdleConns = cfg.MaxOpenConns
	}
	if cfg.ConnMaxIdleTime <= units.Duration(time.Minute) {
		cfg.ConnMaxIdleTime = units.Duration(5 * time.Minute)
	}
	if cfg.ConnMaxLifetime < units.Duration(time.Minute) {
		cfg.ConnMaxLifetime = units.Duration(5 * time.Minute)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime))
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime))

	return db, nil
}

type loggerAdaptor struct {
	slowThreshold time.Duration
}

func (la *loggerAdaptor) LogMode(logger.LogLevel) logger.Interface {
	return &loggerAdaptor{slowThreshold: la.slowThreshold}
}

func (*loggerAdaptor) Info(ctx context.Context, fmt string, args ...interface{}) {
	log.With("module", "gorm").AddCallerSkip(1).CtxInfof(ctx, fmt, args...)
}

func (*loggerAdaptor) Warn(ctx context.Context, fmt string, args ...interface{}) {
	log.With("module", "gorm").AddCallerSkip(1).CtxWarnf(ctx, fmt, args...)
}

func (*loggerAdaptor) Error(ctx context.Context, fmt string, args ...interface{}) {
	log.With("module", "gorm").AddCallerSkip(1).CtxErrorf(ctx, fmt, args...)
}
func (la *loggerAdaptor) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil:
		// ignore not found
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		strSql, rows := fc()
		log.CtxErrorw(ctx, "error sql", "err", err, "elapsed", elapsed, "sql", strSql, "rows", rows)
	case elapsed > la.slowThreshold && la.slowThreshold != 0:
		strSql, rows := fc()
		log.CtxWarnw(ctx, "slow sql", "elapsed", elapsed, "sql", strSql, "rows", rows)
	}
}
