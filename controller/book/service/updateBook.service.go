package service

import (
	"database/sql"
	"mysql/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateBookById(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var book model.BookInput
	err := c.BindJSON(&book)
	if err != nil {
		panic(err)
	}
	query := "UPDATE books SET name = ?, year = ?, author = ?, summary = ?, publisher = ?, pageCount = ?, readPage = ?, reading = ? WHERE id = ?;"
	stmt, err := db.PrepareContext(c, query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    err,
			"status": http.StatusBadRequest,
		})
		panic(err)
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(
		c,
		book.Name,
		book.Year,
		book.Author,
		book.Summary,
		book.Publisher,
		book.PageCount,
		book.ReadPage,
		book.Reading,
		id,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    err,
			"status": http.StatusBadRequest,
		})
		panic(err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    err,
			"status": http.StatusBadRequest,
		})
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":           "ok",
		"status":        http.StatusOK,
		"affected-rows": rows,
	})
}
