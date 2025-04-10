package mysql

import (
	"database/sql"
	"fmt"
	"go-payment-api-server/pkg/logger"
	"time"

	"github.com/samber/lo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

const (
	maxRetries    = 3
	retryInterval = 2 * time.Second
)

func NewDB(cfg Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FTokyo",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	var db *gorm.DB
	var sqlDB *sql.DB

	// NOTE: mysql 起動確認・接続
	_, _, err := lo.AttemptWithDelay(maxRetries, retryInterval, func(i int, d time.Duration) error {
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Warn),
		})
		if err != nil {
			return err
		}

		sqlDB, err = db.DB()
		if err != nil {
			return err
		}

		if err := sqlDB.Ping(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		logger.Log.Error("failed to db init", "db-instance-error", err)
		lo.Must0(err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	logger.Log.Info("connected to database", "host", cfg.Host, "port", cfg.Port)
	return db
}
