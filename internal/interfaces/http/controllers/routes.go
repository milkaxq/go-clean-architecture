package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("health-check", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{}) })
	}
}
