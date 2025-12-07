package payments

import (
	"context"
	"os"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

// StripeService provides simple wrappers around Stripe PaymentIntents.
type StripeService struct {
	apiKey string
}

// GetStripeService no longer requires a string param. It reads the key from environment.
// Use fx.Provide(GetStripeService) so fx can construct the service without needing a raw string.
func GetStripeService() *StripeService {
	apiKey := os.Getenv("STRIPE_API_KEY")
	return &StripeService{apiKey: apiKey}
}

// CreatePaymentIntent creates a PaymentIntent and returns its ID.
// amount is in the smallest currency unit (e.g., cents).
func (s *StripeService) CreatePaymentIntent(ctx context.Context, req CreatePaymentRequest) (*stripe.PaymentIntent, error) {
	stripe.Key = s.apiKey

	params := &stripe.PaymentIntentParams{
		Amount:      stripe.Int64(req.Amount),
		Currency:    stripe.String(req.Currency),
		Description: stripe.String(req.Description),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	return pi, nil
}

// GetPaymentIntent retrieves a PaymentIntent by ID.
func (s *StripeService) GetPaymentIntent(ctx context.Context, id string) (*stripe.PaymentIntent, error) {
	stripe.Key = s.apiKey
	return paymentintent.Get(id, nil)
}

// ... znotraj StripeService struct ...

// ConfirmPaymentIntent posodobi in/ali potrdi PaymentIntent.
// To morda ni vedno potrebno, če uporabljate sodobni Stripe element.
func (s *StripeService) ConfirmPaymentIntent(ctx context.Context, id string, paymentMethodID string) (*stripe.PaymentIntent, error) {
	stripe.Key = s.apiKey

	// params lahko vključujejo payment_method, confirmation_method, itd.
	params := &stripe.PaymentIntentParams{
		// Pri potrditvi pogosto potrebujete PaymentMethod ID
		PaymentMethod: stripe.String(paymentMethodID),
		Confirm:       stripe.Bool(true),
	}

	// Za potrditev se pogosto uporabi kar Update, z nastavitvijo Confirm: true
	pi, err := paymentintent.Update(id, params)
	if err != nil {
		return nil, err
	}

	return pi, nil
}
