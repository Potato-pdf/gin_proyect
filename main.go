package main

import (
	"Gin/db"
	models "Gin/src/domain/models"
	"Gin/src/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//**! Database connection
	DBConnection := db.DBConnection()
	if DBConnection == nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database")

	//Migrations
	DBConnection.AutoMigrate(&models.User{}, &models.Ajolote{}, &models.Task{})

	//**! Router
	r := gin.New()
	routes.SetupRouter(r)

	//**! Run server
	r.Run(":3000")

}
