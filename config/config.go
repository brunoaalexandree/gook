package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

// Is common in Go to have a function that returns only an error or nothing.
func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
