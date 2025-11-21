package controller

import (
	types "Gin/src/domain/types"
	types2 "Gin/src/domain/types/ajolotes"
	http "net/http"

	"github.com/gin-gonic/gin"
)

var prefijos []types.Prefijo
var ajolotes []types2.Ajolotes

func CreatePrefijo(c *gin.Context) {
	var nuevoPrefijo types.Prefijo
	if err := c.ShouldBindJSON(&nuevoPrefijo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prefijos = append(prefijos, nuevoPrefijo)
	c.JSON(http.StatusCreated, nuevoPrefijo)
}

func GetPrefijos(c *gin.Context) {
	c.JSON(http.StatusOK, prefijos)
}

func GetContext(c *gin.Context) {
	context := c.Param("context")
	c.String(http.StatusOK, "Que es dios? %s", context)
}

func CreateAjolote(c *gin.Context) {
	var nuevoAjolote types2.Ajolotes
	if err := c.ShouldBindJSON(&nuevoAjolote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ajolotes = append(ajolotes, nuevoAjolote)
	c.JSON(http.StatusCreated, ajolotes)
}
