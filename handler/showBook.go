package handler

import (
	"net/http"

	"github.com/brunoaalexandree/gook/schemas"
	"github.com/gin-gonic/gin"
)

func GetBookHandler(ctx *gin.Context) {
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

	sendSuccess(ctx, "showBook", book)
}
