package handler

import (
	"net/http"

	"github.com/brunoaalexandree/gook/schemas"
	"github.com/gin-gonic/gin"
)

func CreateBookHandler(ctx *gin.Context) {
	request := CreateBookRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	book := schemas.Book{
		Title:           request.Title,
		Author:          request.Author,
		Genre:           request.Genre,
		PublicationDate: request.PublicationDate,
		Status:          request.Status,
		Rating:          request.Rating,
		Pages:           request.Pages,
	}

	if err := db.Create(&book).Error; err != nil {
		logger.ErrF("error creating book: %v", err)
		sendError(ctx, http.StatusInternalServerError, "erro creating book on database")
		return
	}

	sendSuccess(ctx, "createBook", book)
}
