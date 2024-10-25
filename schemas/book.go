package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title           string
	Author          string
	Genre           string
	PublicationDate time.Time
	Status          string
	Rating          int
	Pages           int
}
