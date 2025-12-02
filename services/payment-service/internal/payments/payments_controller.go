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

// GetPaymentsHandler godoc
// @Summary Get all payments
// @Description Returns a list of all payments admaksdmaslkdmaslkdmnaslkdmnasdlkma
// @Tags payments
// @Produce json
// @Success 200 {array} Payment
// @Failure 500 {object} map[string]string
// @Router /payments [get]
func (c *PaymentsController) GetPaymentsHandler(ctx *gin.Context) {
	payments, err := c.service.GetPayments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, payments)
}
