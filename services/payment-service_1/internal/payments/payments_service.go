package payments

import (
	"context"
)

// PaymentsService handles business logic for payments
type PaymentsService struct {
	repo          *PaymentsRepository
	stripeService *StripeService
}

// GetPaymentsService creates a new PaymentsService
func GetPaymentsService(
	repo *PaymentsRepository,
	stripeService *StripeService,
) *PaymentsService {

	return &PaymentsService{
		repo:          repo,
		stripeService: stripeService,
	}
}

// GetPayments returns all payments
func (s *PaymentsService) GetPayments() ([]Payment, error) {
	return s.repo.GetPayments()
}

// CreatePayment creates a Stripe PaymentIntent and optionally stores it in DB
func (s *PaymentsService) CreatePayment(createPaymentRequest CreatePaymentRequest) (string, error) {
	// 1. Stripe PaymentIntent
	paymentIntent, err := s.stripeService.CreatePaymentIntent(
		context.Background(),
		createPaymentRequest,
	)
	if err != nil {
		return "", err
	}

	// 2. Optional: save to DB
	// payment := Payment{StripeID: paymentIntentID, Amount: amount}
	// err = s.repo.SavePayment(payment)

	return paymentIntent.ID, nil
}
