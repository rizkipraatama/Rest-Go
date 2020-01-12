package main

import (
	"./config"
	"./controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/transaction", inDB.CreateTransaction)
	router.POST("/register", inDB.Register)
	router.POST("/login", inDB.Login)
	router.Run(":3001")
}
