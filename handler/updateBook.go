package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateBookHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Book",
	})
}
