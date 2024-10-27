package handler

import (
	"net/http"

	"github.com/brunoaalexandree/gook/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateBookHandler(ctx *gin.Context) {
	request := UpdateBookRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	book := schemas.Book{}

	if err := db.First(&book, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "book not found")
		return
	}

	if request.Title != "" {
		book.Title = request.Title
	}
	if request.Author != "" {
		book.Author = request.Author
	}
	if request.Status != "" {
		book.Status = request.Status
	}
	if request.Rating != nil {
		book.Rating = request.Rating
	}

	if err := db.Save(&book).Error; err != nil {
		logger.ErrF("error updating book: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error updating book")
		return
	}

	sendSuccess(ctx, "updateBook", book)
}
