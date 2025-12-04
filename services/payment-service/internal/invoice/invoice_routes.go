package invoice

import "payment-service/pkg/lib"

// InvoiceRoutes struct
type InvoiceRoutes struct {
	logger            lib.Logger
	router            *lib.Router
	invoiceController *InvoiceController
}

// SetInvoiceRoutes returns a InvoiceRoutes struct
func SetInvoiceRoutes(
	logger lib.Logger,
	router *lib.Router,
	invoiceController *InvoiceController,
) InvoiceRoutes {
	return InvoiceRoutes{
		logger:            logger,
		router:            router,
		invoiceController: invoiceController,
	}
}

// Setup registers the invoice routes
func (route InvoiceRoutes) Setup() {
	route.logger.Info("Setting up [INVOICE] routes.")

	api := route.router.Group("/invoices")
	{
		api.GET("/", route.invoiceController.GetInvoicesHandler)
		api.GET("/:id", route.invoiceController.GetInvoiceByIdHandler)
		api.POST("/", route.invoiceController.CreateInvoiceHandler)
		api.PUT("/:id", route.invoiceController.UpdateInvoiceHandler)
		api.DELETE("/:id", route.invoiceController.DeleteInvoiceHandler)
	}
}
