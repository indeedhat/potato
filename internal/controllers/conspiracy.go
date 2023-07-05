package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controllers) GetConspiracy(ctx *gin.Context) {
	theory, err := c.repo.GetRandom()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Conspiracy theory not found",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"conspiracy": theory.Text,
	})
}
