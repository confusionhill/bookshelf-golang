package service

import (
	"database/sql"
	"mysql/model"
	"mysql/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context, db *sql.DB) {
	var user model.UserLoginModel
	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}
	query := "SELECT * FROM user WHERE username = ?"
	rows, err := db.Query(query, user.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":    "Account not found",
			"status": http.StatusNotFound,
		})
	}
	var userDetail model.UserDetailModel
	if rows.Next() {
		err := rows.Scan(
			&userDetail.Uuid,
			&userDetail.Username,
			&userDetail.Password,
			&userDetail.Email,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":    err,
				"status": http.StatusBadRequest,
			})
			panic(err)
		}
	}
	token, err := token.GenerateToken(userDetail.Uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    "Token could not be generated",
			"status": http.StatusBadRequest,
		})
		panic(err)
	}
	if CheckPasswordHash(user.Password, userDetail.Password) {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "ok",
			"status": http.StatusOK,
			"token":  token,
		})
	}
	defer rows.Close()

}
