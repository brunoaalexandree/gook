package handler

import (
	"fmt"
	"time"
)

func errParamIsRequired(name string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateBookRequest struct {
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	Genre           string    `json:"genre"`
	PublicationDate time.Time `json:"publication_date"`
	Status          string    `json:"status"`
	Rating          *int      `json:"rating"`
	Pages           int       `json:"pages"`
}

func (r *CreateBookRequest) Validate() error {
	if r.Title == "" && r.Author == "" && r.Rating == nil && r.Pages <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.Author == "" {
		return errParamIsRequired("author", "string")
	}
	if r.Status == "" {
		return errParamIsRequired("status", "string")
	}
	if r.Rating == nil {
		return errParamIsRequired("rating", "int")
	}
	if r.Pages <= 0 {
		return errParamIsRequired("pages", "int")
	}

	return nil
}

type UpdateBookRequest struct {
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	Genre           string    `json:"genre"`
	PublicationDate time.Time `json:"publication_date"`
	Status          string    `json:"status"`
	Rating          *int      `json:"rating"`
	Pages           int       `json:"pages"`
}

func (r *UpdateBookRequest) Validate() error {
	if r.Title != "" || r.Author != "" || r.Rating != nil || r.Pages <= 0 {
		return nil
	}

	return fmt.Errorf("at least one field must be provided")
}
