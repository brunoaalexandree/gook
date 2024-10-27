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
	Rating          *int
	Pages           int
}

type BookResponse struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at,omitempty"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	Genre           string    `json:"genre"`
	PublicationDate time.Time `json:"publication_date"`
	Status          string    `json:"status"`
	Rating          *int      `json:"rating"`
	Pages           int       `json:"pages"`
}
