package payments

import (
	"payment-service/pkg/lib"
)

// PaymentsRoutes struct
type PaymentsRoutes struct {
	logger             lib.Logger
	router             *lib.Router
	paymentsController *PaymentsController
}

// SetPaymentsRoutes returns a PaymentsRoutes struct
func SetPaymentsRoutes(
	logger lib.Logger,
	router *lib.Router,
	paymentsController *PaymentsController,
) PaymentsRoutes {
	return PaymentsRoutes{
		logger:             logger,
		router:             router,
		paymentsController: paymentsController,
	}
}

// Setup registers the payment routes
func (route PaymentsRoutes) Setup() {
	route.logger.Info("Setting up [PAYMENTS] routes.")

	// Register /payments route
	api := route.router.Group("/payments")
	{
		api.GET("/", route.paymentsController.GetPaymentsHandler)
		// kasneje lahko dodamo POST, PUT, DELETE
	}
}
