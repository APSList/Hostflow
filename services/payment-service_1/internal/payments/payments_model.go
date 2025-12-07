package payments

import "time"

type Payment struct {
	Id              string    `json:"id" db:"id"`
	OrganizationId  string    `json:"organization_id" db:"organization_id"`
	ReservationId   string    `json:"reservation_id" db:"reservation_id"`
	Amount          float64   `json:"amount" db:"amount"`
	PaymentIntentId string    `json:"payment_intent_id" db:"payment_intent_id"`
	PaymentMethod   string    `json:"payment_method" db:"payment_method"`
	Status          string    `json:"status" db:"status"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// CreatePaymentRequest defines the request body for creating a payment
type CreatePaymentRequest struct {
	Amount        int64   `json:"amount"`
	Currency      string  `json:"currency"`
	Description   string  `json:"description"`
	PaymentMethod *string `json:"payment_method,omitempty"`
}

// CreatePaymentResponse defines the successful response
type CreatePaymentResponse struct {
	PaymentIntentId string `json:"payment_intent_id"`
}
