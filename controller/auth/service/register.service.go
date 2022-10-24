package service

import (
	"database/sql"
	"log"
	"mysql/model"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func RegisterUser(c *gin.Context, db *sql.DB) {
	var user model.UserModel
	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}
	if !isEmailValid(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "Wrong email format",
		})
		return
	}
	password, err := HashPassword(user.Password)
	if err != nil {
		panic(err)
	}
	query := "INSERT INTO user (username, password, email) VALUES(?,?,?)"
	stmt, err := db.PrepareContext(c, query)
	if err != nil {
		panic(err)
	}
	res, err := stmt.ExecContext(
		c,
		user.Username,
		password,
		user.Email,
	)
	if err != nil {
		panic(err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
	}
	if rows >= 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg":           "ok",
			"status":        http.StatusOK,
			"affected-rows": rows,
		})
	}
}
