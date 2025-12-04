package payments

import (
	"go.uber.org/fx"
)

// ======== EXPORTS ========

// Module exports services present
var Context = fx.Options(
	fx.Provide(GetPaymentsController),
	fx.Provide(GetPaymentsService),
	fx.Provide(GetPaymentsRepository),
	fx.Provide(SetPaymentsRoutes),
	fx.Provide(GetStripeService),
)
