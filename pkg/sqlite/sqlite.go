package sqlite

import (
	"fmt"
	"sync"

	"github.com/VQIVS/web3-tracker.git/internal/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

var (
	instance *DB
	once     sync.Once
)

func SetupDatabase(databasePath string) (*DB, error) {
	db, err := newConnection(databasePath)
	if err != nil {
		return nil, err
	}

	if err := migrate(db, &entities.Wallet{}); err != nil {
		return nil, err
	}

	return db, nil
}

func newConnection(dbPath string) (*DB, error) {
	var err error
	once.Do(func() {
		db, dbErr := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if dbErr != nil {
			err = fmt.Errorf("failed to open database: %v", dbErr)
			return
		}

		sqlDB, dbErr := db.DB()
		if dbErr != nil {
			err = fmt.Errorf("failed to get underlying database: %v", dbErr)
			return
		}

		if dbErr := sqlDB.Ping(); dbErr != nil {
			err = fmt.Errorf("failed to ping database: %v", dbErr)
			return
		}
		instance = &DB{db}
	})
	return instance, err
}
func migrate(db *DB, models ...interface{}) error {
	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}
	return nil
}
