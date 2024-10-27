package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

// Is common in Go to have a function that returns only an error or nothing.
func Init() error {
	var err error

	db, err = InitializePostgres()
	if err != nil {
		return fmt.Errorf("🛑 error initializing postgres: %v", err)
	}

	return nil
}

func GetProstgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
