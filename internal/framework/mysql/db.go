// Package db implements BookRepoPort with mysql
package db

import (
	"fmt"
	"hexagonal-arch/helper/errors"
	"hexagonal-arch/internal/ports/types"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"hexagonal-arch/config"
	"hexagonal-arch/internal/ports"
)

// Adapter implements ports.BookRepoPort with mysql
type Adapter struct {
	db *gorm.DB
}

// This validates that Adapter implements the ports.BookRepoPort interface
var _ ports.BookRepoPort = Adapter{}

// NewAdapterWithConfig creates a new adapter with the given config
func NewAdapterWithConfig(cfg config.DBConfig) (*Adapter, errors.Error) {
	var adapter *Adapter
	var cErr errors.Error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	dialect := mysql.Open(dsn)
	adapter, cErr = newAdapterWithDialector(dialect)
	if cErr != nil {
		return nil, cErr
	}

	sqlDB, err := adapter.db.DB()
	if err != nil {
		return nil, errors.InternalError(err)
	}

	sqlDB.SetConnMaxLifetime(cfg.MaxConnTime)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	return adapter, nil
}

// newAdapterWithDialector creates a new adapter with the given dialect
func newAdapterWithDialector(dialect gorm.Dialector) (*Adapter, errors.Error) {
	d, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now()
		},
	})
	if err != nil {
		return nil, errors.InternalError(err)
	}

	adapter := &Adapter{
		db: d,
	}

	createCallback := d.Callback().Create()
	createCallback.Clauses = []string{"INSERT", "VALUES"}

	return adapter, nil
}

func (a Adapter) migrate() errors.Error {
	db := a.db.Session(&gorm.Session{
		Logger: logger.Default.LogMode(logger.Error),
	})

	err := db.AutoMigrate(&types.Book{})
	if err != nil {
		return errors.InternalDBError(err)
	}

	return nil
}

// BeginTx starts a new DB transaction
func (a Adapter) BeginTx() ports.BookRepoPort {
	return Adapter{
		db: a.db.Begin(),
	}
}

// RollbackTx rolls back the current DB transaction
func (a Adapter) RollbackTx() {
	a.db.Rollback()
}

// CommitTx commits the current DB transaction
func (a Adapter) CommitTx() errors.Error {
	if err := a.db.Commit().Error; err != nil {
		return errors.InternalDBError(err)
	}
	return nil
}
