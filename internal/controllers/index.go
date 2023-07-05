package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controllers) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
