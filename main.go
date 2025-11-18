package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
) 

func main (){

	r := gin.New()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"mensaje":"hola",
		})
	})

	r.GET("/saludo/:nombre", func(ctx *gin.Context) {
		nombre := ctx.Param("nombre")
		ctx.String(http.StatusOK, "Hola, %s", nombre)
	})

	r.Run(":8000")

}