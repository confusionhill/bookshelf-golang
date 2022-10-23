package main

import (
	"mysql/controller"
	"mysql/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.CreateConnection()
	r := gin.Default()
	controller.BookController(r, db)
	r.Run(":3030") // listen and serve on 0.0.0.0:3030 (for windows "localhost:8080")
	defer db.Close()
}
