package routes

import (
	"net/http"

	"Gin/src/domain/types"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	RoutesGroup := r.Group("/prefijo")
	RoutesGroup.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"SAS": "Down",
		})
	})
	RoutesGroup.GET("/spotify/:context", func(ctx *gin.Context) {
		nombre := ctx.Param("context")
		ctx.String(http.StatusOK, "Que es dios? %s", nombre)
	})

	RoutesGroup.POST("/create", func(ctx *gin.Context) {
		var nuevoPrefijo types.Prefijo
		if err := ctx.ShouldBindJSON(&nuevoPrefijo); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Recibido",
			"data":    nuevoPrefijo,
		})
	})
}
