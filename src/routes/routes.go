package routes

import (
	"Gin/src/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	RoutesGroup := r.Group("/prefijo")
	RoutesGroup.GET("/", controller.GetPrefijos)
	RoutesGroup.GET("/spotify/:context", controller.GetContext)
	RoutesGroup.POST("/create", controller.CreatePrefijo)
	RoutesGroup.POST("/ajolote", controller.CreateAjolote)
}
kk