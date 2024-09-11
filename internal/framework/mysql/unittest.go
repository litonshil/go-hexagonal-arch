package db

import (
	"database/sql"
	"fmt"
	"hexagonal-arch/config"
	"hexagonal-arch/helper/errors"
	"os"
	"strings"
	"testing"
)

// NewAdapterForTest creates a new adapter for testing
func NewAdapterForTest(t *testing.T) (*Adapter, errors.Error) {
	err := os.Setenv("TZ", "UTC")
	if err != nil {
		return nil, errors.InternalError(err)
	}

	config.Init()

	cfg := config.GetDBConfig()

	cfg.Name = fmt.Sprintf("%s_%s", cfg.Name, strings.ToLower(t.Name()))
	fmt.Printf("Running Test with MySQL (%s)\n", cfg.Name)
	if cErr := createMySQLDatabase(cfg); cErr != nil {
		return nil, cErr
	}

	// To run the tests in parallel, we need to create a new database for each test
	t.Parallel()

	// Cleanup database after the test
	t.Cleanup(func() {
		if cErr := deleteMySQLDatabase(cfg); cErr != nil {
			fmt.Println("Failed to delete database", err)
			return
		}
	})

	adapter, cErr := NewAdapterWithConfig(cfg)
	if err != nil {
		return nil, cErr
	}

	if cErr = adapter.migrate(); err != nil {
		return nil, cErr
	}

	return adapter, nil
}

func createMySQLDatabase(cfg config.DBConfig) errors.Error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	_, err = db.Exec("DROP DATABASE IF EXISTS " + cfg.Name)
	if err != nil {
		return errors.InternalDBError(err)
	}

	_, err = db.Exec("CREATE DATABASE " + cfg.Name)
	if err != nil {
		return errors.InternalDBError(err)
	}

	return nil
}

func deleteMySQLDatabase(cfg config.DBConfig) errors.Error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	_, err = db.Exec("DROP DATABASE " + cfg.Name)
	if err != nil {
		return errors.InternalDBError(err)
	}
	return nil
}
