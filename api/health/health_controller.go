package health

import (
	"net/http"
	response "urlshortener/utils"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	healthService *HealthService
}

func NewHealthController(
	healthService *HealthService,
) *HealthController {
	return &HealthController{
		healthService: healthService,
	}
}

func (ctx *HealthController) HealthCheck(c *gin.Context) {
	err := ctx.healthService.HealthCheck()
	if err != nil {
		response.Error(c, "Health check failed")
		return
	}

	// response.Message(c, "API running sucessfully.", http.StatusOK)
	c.JSON(http.StatusOK, gin.H{
		"message": "API running sucessfully.",
		"status":  http.StatusOK,
	})
}
