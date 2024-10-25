package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBookHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Book",
	})
}
