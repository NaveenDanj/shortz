package db

import (
	"config"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

func ConnectPostgres(dsn string) (*Database, error) {
	gormConfig := &gorm.Config{
		Logger: logger.New(
			log.Default(),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      false,
			},
		),
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(config.PGPoolConsts.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(config.PGPoolConsts.MaxIdleConnections)
	sqlDB.SetConnMaxLifetime(time.Duration(config.PGPoolConsts.ConnectionMaxLifeTime) * time.Minute)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.PGPoolConsts.ConnectionMaxIdleTime) * time.Minute)

	return &Database{DB: db}, nil
}
