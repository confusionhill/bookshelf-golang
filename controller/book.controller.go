package controller

import (
	"database/sql"
	"log"
	"mysql/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func BookController(
	router *gin.Engine,
	db *sql.DB,
) {
	bookRouter := router.Group("/books")
	bookRouter.POST("/", func(c *gin.Context) {
		AddBook(c, db)
	})
	bookRouter.GET("/", func(c *gin.Context) {
		GetBook(c, db)
	})
	bookRouter.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		GetBookById(c, db, id)
	})

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
	var book model.BookInput
	if rows.Next() {
		err := rows.Scan(
			&book.Id,
			&book.Name,
			&book.Year,
			&book.Author,
			&book.Summary,
			&book.Publisher,
			&book.PageCount,
			&book.PageCount,
			&book.ReadPage,
			&book.Reading)
		if err != nil {
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "ok",
		"status": http.StatusOK,
		"data":   book,
	})
}

func GetBook(c *gin.Context, db *sql.DB) {
	query := "SELECT * FROM books"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    err,
			"status": http.StatusBadRequest,
		})
	}
	listOfBooks := []model.BookInput{}
	for rows.Next() {
		var book model.BookInput
		err := rows.Scan(
			&book.Id,
			&book.Name,
			&book.Year,
			&book.Author,
			&book.Summary,
			&book.Publisher,
			&book.PageCount,
			&book.PageCount,
			&book.ReadPage,
			&book.Reading)
		if err != nil {
			panic(err)
		}
		listOfBooks = append(listOfBooks, book)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "done",
		"status": http.StatusOK,
		"data":   listOfBooks,
	})
}

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
		log.Printf("Error %s when preparing SQL statement", err)
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

func turnBoolToStr(val bool) string {
	if val {
		return "done"
	}
	return "false"
}
