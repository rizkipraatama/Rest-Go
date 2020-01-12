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

	err := idb.DB.Where(map[string]interface{}{"Username": username, "Password": password}).First(&Login).Error
	if err == nil {
		result = gin.H{
			"result": Login,
		}
	} else {
		result = gin.H{
			"result": "Username/Password Doesn't Match",
		}
	}
	c.JSON(http.StatusOK, result)
}
