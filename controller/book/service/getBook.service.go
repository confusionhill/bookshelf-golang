package service

import (
	"database/sql"
	"mysql/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context, db *sql.DB) {
	query := "SELECT * FROM books"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    err,
			"status": http.StatusBadRequest,
		})
	}
	listOfBooks := []model.BookModel{}
	for rows.Next() {
		var book model.BookModel
		err := rows.Scan(
			&book.Id,
			&book.Name,
			&book.Year,
			&book.Author,
			&book.Summary,
			&book.Publisher,
			&book.PageCount,
			&book.ReadPage,
			&book.Reading,
			&book.Finished,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":    err,
				"status": http.StatusBadRequest,
			})
			panic(err)
		}
		listOfBooks = append(listOfBooks, book)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "done",
		"status": http.StatusOK,
		"data":   listOfBooks,
	})
	defer rows.Close()
}

func GetBookById(c *gin.Context, db *sql.DB, id string) {
	query := "SELECT * FROM books WHERE id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    err,
			"status": http.StatusBadRequest,
		})
	}
	var book model.BookModel
	if rows.Next() {
		err := rows.Scan(
			&book.Id,
			&book.Name,
			&book.Year,
			&book.Author,
			&book.Summary,
			&book.Publisher,
			&book.PageCount,
			&book.ReadPage,
			&book.Reading,
			&book.Finished,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":    err,
				"status": http.StatusBadRequest,
			})
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "ok",
		"status": http.StatusOK,
		"data":   book,
	})
	defer rows.Close()
}
