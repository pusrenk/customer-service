package database

import (
	"fmt"

	"github.com/pusrenk/customer-service/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDatabase init database
func InitDatabase(cfg *configs.Config) (*gorm.DB, error) {
	loggerMode := logger.Silent
	if cfg.App.Env != "production" {
		loggerMode = logger.Info
	}

	db, err := gorm.Open(
		postgres.Open(connectionString(cfg)),
		&gorm.Config{
			Logger: logger.Default.LogMode(loggerMode),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to start Gorm DB: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxIdleTime(cfg.Database.ConnMaxIdleTime)
	sqlDB.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)

	return db, nil
}

// connectionString generate connection dsn string
func connectionString(cfg *configs.Config) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Database,
		cfg.Database.Password,
	)
}
