package main

import (
	"context"
	"mysql/controller"
	"mysql/database"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.CreateConnection()
	r := gin.Default()
	ctx := context.Background()
	controller.BookController(r, db)
	r.GET("/ping", func(c *gin.Context) {
		rows, err := db.QueryContext(ctx, "SELECT * FROM customer;")
		if err != nil {
			panic(err)
		}
		var id, name string
		for rows.Next() {
			qErr := rows.Scan(&id, &name)
			if qErr != nil {
				panic(qErr)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"id":   id,
			"name": name,
		})
	})
	r.Run(":3030") // listen and serve on 0.0.0.0:3030 (for windows "localhost:8080")
	defer db.Close()
}
