package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine{
	r := gin.New()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"SAS":"Down",
	})
	})
	r.GET("/spotify/:contex", func(ctx *gin.Context) {
		nombre := ctx.Param("context")
		ctx.String(http.StatusOK, "Que es dios? %s", nombre)
	})
	return r
}