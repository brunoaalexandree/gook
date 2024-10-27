package handler

import (
	"net/http"

	"github.com/brunoaalexandree/gook/schemas"
	"github.com/gin-gonic/gin"
)

func GetBooksHandler(ctx *gin.Context) {
	books := []schemas.Book{}

	if err := db.Find(&books).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing books")
	}

	sendSuccess(ctx, "listBooks", books)
}
