package controllers

import "github.com/gin-gonic/gin"

type HealthController interface {
	HealthCheck(c *gin.Context)
}

type healthController struct {
}

func NewHealthController() HealthController {
	return &healthController{}
}

func (h *healthController) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
