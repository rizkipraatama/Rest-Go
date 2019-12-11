package controllers

import (
	"net/http"

	"../structs"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) Register(c *gin.Context) {
	var (
		Register structs.Register
		result   gin.H
	)
	firstname := c.PostForm("first_name")
	lastname := c.PostForm("last_name")
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirmpassword := c.PostForm("confirm_password")
	email := c.PostForm("email")
	phonenumber := c.PostForm("phone_number")
	err := idb.DB.Where("username = ?", username).First(&Register).Error
	if err == nil {
		result = gin.H{
			"result": "Username is already taken",
		}
	} else {
		if password != confirmpassword {
			result = gin.H{
				"result": "Password Doesn't Match",
			}
		} else {
			Register.First_Name = firstname
			Register.Last_Name = lastname
			Register.Username = username
			Register.Password = password
			Register.Email = email
			Register.Phone_Number = phonenumber
			idb.DB.Create(&Register)
			result = gin.H{
				"result": Register,
			}
		}
	}
	c.JSON(http.StatusOK, result)
}
