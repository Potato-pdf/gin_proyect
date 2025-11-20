package main

import (
	"Gin/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()
	routes.SetupRouter(r)
	r.Run(":3000")

}
