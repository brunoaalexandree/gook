package config

import (
	"fmt"
	"os"

	"github.com/brunoaalexandree/gook/schemas"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	err := godotenv.Load()
	if err != nil {
		logger.ErrF("🏠 environment variables load error: %v", err)
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrF("Postgres initialization error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Book{})
	if err != nil {
		logger.ErrF("Postgres automigration error: %v", err)
		return nil, err
	}
	return db, nil
}
