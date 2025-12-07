package invoice

import (
	"net/http"
	"payment-service/pkg/models"

	"github.com/gin-gonic/gin"
)

// You can keep this alias for use inside the Go code logic,
// but we will use 'models.ErrorResponse' in the comments below.
type ErrorResponse = models.ErrorResponse

type InvoiceController struct {
	service *InvoiceService
}

func GetInvoiceController(service *InvoiceService) *InvoiceController {
	return &InvoiceController{service: service}
}

// GetInvoicesHandler godoc
// @Summary      Get all invoices
// @Description  Returns a list of all invoices
// @Tags         invoices
// @Produce      json
// @Success      200  {array}   invoice.Invoice
// @Failure      500  {object}  models.ErrorResponse
// @Router       /invoices [get]
func (c *InvoiceController) GetInvoicesHandler(ctx *gin.Context) {
	invoices, err := c.service.GetInvoices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, invoices)
}

// GetInvoiceByIdHandler godoc
// @Summary      Get invoice by ID
// @Description  Returns an invoice based on ID
// @Tags         invoices
// @Produce      json
// @Param        id   path      string  true  "Invoice ID"
// @Success      200  {object}  invoice.Invoice
// @Failure      404  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /invoices/{id} [get]
func (c *InvoiceController) GetInvoiceByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	invoice, err := c.service.GetInvoiceById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, invoice)
}

// CreateInvoiceHandler godoc
// @Summary      Create a new invoice
// @Description  Creates a new invoice record
// @Tags         invoices
// @Accept       json
// @Produce      json
// @Param        invoice  body      invoice.CreateInvoiceRequest  true  "Invoice info"
// @Success      201      {object}  invoice.Invoice
// @Failure      400      {object}  models.ErrorResponse
// @Failure      500      {object}  models.ErrorResponse
// @Router       /invoices [post]
func (c *InvoiceController) CreateInvoiceHandler(ctx *gin.Context) {
	var req CreateInvoiceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	invoice, err := c.service.CreateInvoice(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, invoice)
}

// UpdateInvoiceHandler godoc
// @Summary      Update an invoice
// @Description  Updates an existing invoice
// @Tags         invoices
// @Accept       json
// @Produce      json
// @Param        id       path      string                        true  "Invoice ID"
// @Param        invoice  body      invoice.UpdateInvoiceRequest  true  "Invoice data"
// @Success      200      {object}  invoice.Invoice
// @Failure      400      {object}  models.ErrorResponse
// @Failure      404      {object}  models.ErrorResponse
// @Failure      500      {object}  models.ErrorResponse
// @Router       /invoices/{id} [put]
func (c *InvoiceController) UpdateInvoiceHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var req UpdateInvoiceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	invoice, err := c.service.UpdateInvoice(id, req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, invoice)
}

// DeleteInvoiceHandler godoc
// @Summary      Delete invoice
// @Description  Deletes an invoice by ID
// @Tags         invoices
// @Produce      json
// @Param        id   path      string  true  "Invoice ID"
// @Success      200  {object}  map[string]bool
// @Failure      404  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /invoices/{id} [delete]
func (c *InvoiceController) DeleteInvoiceHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.service.DeleteInvoice(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"deleted": true})
}
