package invoice

import "time"

type Invoice struct {
	Id             string    `json:"id" db:"id"`
	OrganizationId string    `json:"organization_id" db:"organization_id"`
	PaymentId      *string   `json:"payment_id" db:"payment_id"` // nullable FK
	InvoiceNumber  string    `json:"invoice_number" db:"invoice_number"`
	CustomerId     string    `json:"customer_id" db:"customer_id"`
	IssueDate      time.Time `json:"issue_date" db:"issue_date"`
	DueDate        time.Time `json:"due_date" db:"due_date"`
	Amount         float64   `json:"amount" db:"amount"`
	TxtAmount      string    `json:"txt_amount" db:"txt_amount"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type CreateInvoiceRequest struct {
	OrganizationId string  `json:"organization_id"`
	PaymentId      *string `json:"payment_id,omitempty"`
	InvoiceNumber  string  `json:"invoice_number"`
	CustomerId     string  `json:"customer_id"`
	IssueDate      string  `json:"issue_date"` // ISO string, will parse in service
	DueDate        string  `json:"due_date"`
	Amount         float64 `json:"amount"`
	TxtAmount      string  `json:"txt_amount"`
}

type UpdateInvoiceRequest struct {
	InvoiceNumber *string  `json:"invoice_number,omitempty"`
	PaymentId     *string  `json:"payment_id,omitempty"`
	CustomerId    *string  `json:"customer_id,omitempty"`
	IssueDate     *string  `json:"issue_date,omitempty"`
	DueDate       *string  `json:"due_date,omitempty"`
	Amount        *float64 `json:"amount,omitempty"`
	TxtAmount     *string  `json:"txt_amount,omitempty"`
}
