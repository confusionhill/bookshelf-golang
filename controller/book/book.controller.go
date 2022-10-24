package book

import (
	"database/sql"
	"mysql/controller/book/service"

	"github.com/gin-gonic/gin"
)

func BookController(
	router *gin.Engine,
	db *sql.DB,
) {
	bookRouter := router.Group("/books")
	bookRouter.POST("/", func(c *gin.Context) {
		service.AddBook(c, db)
	})
	bookRouter.GET("/", func(c *gin.Context) {
		service.GetBook(c, db)
	})
	bookRouter.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		service.GetBookById(c, db, id)
	})
	bookRouter.PUT("/:id", func(c *gin.Context) {
		service.UpdateBookById(c, db)
	})
	bookRouter.DELETE("/:id", func(c *gin.Context) {
		service.DeleteBookById(c, db)
	})
}
