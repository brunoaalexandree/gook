package router

import (
	"github.com/brunoaalexandree/gook/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/book", handler.GetBookHandler)
		v1.POST("/book", handler.CreateBookHandler)
		v1.PUT("/book", handler.UpdateBookHandler)
		v1.DELETE("/book", handler.DeleteBookHandler)
		v1.GET("/books", handler.GetBooksHandler)
	}
}
