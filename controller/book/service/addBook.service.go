package service

import (
	"database/sql"
	"log"
	"mysql/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddBook(c *gin.Context, db *sql.DB) {
	var book model.BookInput
	err := c.BindJSON(&book)
	id := uuid.New()
	if err != nil {
		panic(err)
	}
	query := "INSERT INTO books (id, name, year, author, summary, publisher, pageCount, readPage, reading, finished) VALUES (?,?,?,?,?,?,?,?,?,?);"
	stmt, err := db.PrepareContext(c, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL", err)
		panic(err)
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(
		c,
		id.String(),
		book.Name,
		book.Year,
		book.Author,
		book.Summary,
		book.Publisher,
		book.PageCount,
		book.ReadPage,
		book.Reading,
		false)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    err,
			"status": http.StatusBadRequest,
		})
		panic(err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":           "ok",
		"status":        http.StatusOK,
		"affected-rows": rows,
	})
}
