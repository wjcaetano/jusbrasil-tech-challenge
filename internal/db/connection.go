package db

import (
	"fmt"
	"jusbrasil-tech-challenge/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(cfg config.Configuration) (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Cluster,
		cfg.Database.Name,
	)
	dbConfig := gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(mysql.Open(dns), &dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to configure SQL DB: %w", err)
	}
	sqlDB.SetMaxOpenConns(cfg.HTTPClient.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.HTTPClient.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.HTTPClient.ConnMaxLifetime)

	return db, nil
}
