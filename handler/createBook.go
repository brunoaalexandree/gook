package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBookHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Book",
	})
}
