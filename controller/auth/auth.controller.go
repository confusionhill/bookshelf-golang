package auth

import (
	"database/sql"
	"mysql/controller/auth/service"

	"github.com/gin-gonic/gin"
)

func AuthController(
	router *gin.Engine,
	db *sql.DB,
) {
	router.POST("/register", func(c *gin.Context) {
		service.RegisterUser(c, db)
	})
	router.POST("/login", func(c *gin.Context) {
		service.LoginUser(c, db)
	})
}
