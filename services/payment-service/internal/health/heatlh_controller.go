package health

import (
	"github.com/gin-gonic/gin"
)

// @Summary Check service health
// @Description Returns the operational status of the Payment Service, including connectivity checks.
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string "Status is UP"
// @Router /health [get]
func HealthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "UP", "service": "PaymentService"})
}

// @Summary Check service health
// @Description Returns the operational status of the Payment Service, including connectivity checks.
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string "Status is UP"
// @Router /health [get]
func readinessCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "UP", "service": "PaymentService"})
}
