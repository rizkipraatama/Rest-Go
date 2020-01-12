package controllers

import (
	"net/http"

	"../structs"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) Login(c *gin.Context) {
	var (
		Login  structs.Login
		result gin.H
	)

	username := c.PostForm("username")
	password := c.PostForm("password")

	notFound := idb.DB.Where("Username = ? AND Password = ?", username, password).Find(&Login).RecordNotFound()
	if !notFound {
		result = gin.H{
			"result": Login,
		}
	} else {
		result = gin.H{
			"result": "Not Found",
		}
	}
	c.JSON(http.StatusOK, result)
}
