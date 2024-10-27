package handler

import (
	"fmt"
	"net/http"

	"github.com/brunoaalexandree/gook/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteBookHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	book := schemas.Book{}

	if err := db.First(&book, id).Error; err != nil {
		logger.ErrF("error finding book: %v", err)
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %v not found", id))
		return
	}

	if err := db.Delete(&book, id).Error; err != nil {
		logger.ErrF("error deleting book: %v", err)
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("erro deleting book with id: %v", id))
		return
	}

	sendSuccess(ctx, "deleteBook", book)
}
