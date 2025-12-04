package invoice

import "go.uber.org/fx"

var Context = fx.Options(
	fx.Provide(GetInvoiceController),
	fx.Provide(GetInvoiceService),
	fx.Provide(GetInvoiceRepository),
	fx.Provide(SetInvoiceRoutes),
)
