package service

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteBookById(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	query := "DELETE FROM books WHERE id = ?;"
	stmt, err := db.PrepareContext(c, query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(c, id)
	if err != nil {
		panic(err)
	}
	rows, err := res.RowsAffected()
	if err != nil || rows <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":    err,
			"status": http.StatusNotFound,
		})
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":           "Buku berhasil dihapus",
		"status":        http.StatusOK,
		"affected-rows": rows,
	})

}
