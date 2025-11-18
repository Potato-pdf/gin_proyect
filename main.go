package main 

import("github.com/gin-gonic/gin") 

func main (){

	r := gin.New()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"mensaje":"hola",
		})
	})

	r.Run(":8000")

}