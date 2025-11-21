package main

import (
	"Gin/db"
	"Gin/src/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	DBConnection := db.DBConnection()
	if DBConnection == nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database")
	r := gin.New()
	routes.SetupRouter(r)
	r.Run(":3000")

}
