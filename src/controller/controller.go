package controller

import (
	"Gin/src/domain/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

var prefijos []types.Prefijo

func CreatePrefijo(c *gin.Context) {
	var nuevoPrefijo types.Prefijo
	if err := c.ShouldBindJSON(&nuevoPrefijo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prefijos = append(prefijos, nuevoPrefijo)
	c.JSON(http.StatusCreated, nuevoPrefijo)
}
