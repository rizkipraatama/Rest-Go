package controllers

import (
	"net/http"

	"../structs"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) getTransaction(c *gin.Context) {
	var (
		transaction structs.Transaction
		result      gin.H
	)

	place := c.Param("place")
	err := idb.DB.Where("place = ?", place).First(&transaction).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": transaction,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateTransaction(c *gin.Context) {
	var (
		transaction structs.Transaction
		result      gin.H
	)
	place := c.PostForm("Place")
	total := c.PostForm("Total")
	transaction.Place = place
	transaction.Total = total
	idb.DB.Create(&transaction)
	result = gin.H{
		"result": transaction,
	}
	c.JSON(http.StatusOK, result)
}
