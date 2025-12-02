package payments

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PaymentsController handles HTTP requests
type PaymentsController struct {
	service *PaymentsService
}

// NewPaymentsController creates a new controller
func GetPaymentsController(service *PaymentsService) *PaymentsController {
	return &PaymentsController{
		service: service,
	}
}

// GetPaymentsHandler handles GET /payments
func (c *PaymentsController) GetPaymentsHandler(ctx *gin.Context) {
	payments, err := c.service.GetPayments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payments)
}
