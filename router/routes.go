package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/book", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Book",
			})
		})
		v1.POST("/book", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST Book",
			})
		})
		v1.PUT("/book", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PUT Book",
			})
		})
		v1.DELETE("/book", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "DELETE Book",
			})
		})
		v1.GET("/books", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Books",
			})
		})
	}
}
