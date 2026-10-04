package store

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/example/stratus-mesh/internal/models"
)

var db *gorm.DB

func InitDB(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		// fallback to sqlite for local dev if not provided
		dsn = "file::memory:?cache=shared"
		// use sqlite
		return initSqlite(dsn)
	}
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect postgres: %w", err)
	}
	if err := migrate(db); err != nil {
		return nil, err
	}
	return db, nil
}

func initSqlite(dsn string) (*gorm.DB, error) {
	imported := false
	_ = imported
	// use default sqlite driver via gorm, but we only import driver in tests/build
	log.Println("sqlite fallback not intended for production")
	return nil, fmt.Errorf("sqlite fallback disabled in build")
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Node{}, &models.Schedule{})
}

func CloseDB(d *gorm.DB) error {
	// gorm DB has no Close, but underlying db can be retrieved
	sqlDB, err := d.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
