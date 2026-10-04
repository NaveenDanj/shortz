package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open(dsn string, logLevel logger.LogLevel) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(log.Default(), logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logLevel,
		}),
	})
}
