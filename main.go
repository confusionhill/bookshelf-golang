package main

import (
	"database/sql"
	"mysql/controller/auth"
	"mysql/controller/book"
	"mysql/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.CreateConnection()
	setupServer(db)
	defer db.Close()
}

func setupServer(db *sql.DB) {
	r := gin.Default()
	book.BookController(r, db)
	auth.AuthController(r, db)
	r.Run(":3030") // listen and serve on 0.0.0.0:3030 (for windows "localhost:8080")
}
