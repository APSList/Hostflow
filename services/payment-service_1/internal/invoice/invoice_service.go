package invoice

import (
	"errors"
	"time"
)

// InvoiceService handles business logic for invoices
type InvoiceService struct {
	repo *InvoiceRepository // will be implemented later
}

// Dependency injection constructor
func GetInvoiceService(repo *InvoiceRepository) *InvoiceService {
	return &InvoiceService{
		repo: repo,
	}
}

// GetInvoices returns all invoices
func (s *InvoiceService) GetInvoices() ([]Invoice, error) {
	return s.repo.GetInvoices()
}

// GetInvoiceById returns a single invoice by ID
func (s *InvoiceService) GetInvoiceById(id string) (*Invoice, error) {
	return s.repo.GetInvoiceById(id)
}

// CreateInvoice creates a new invoice
func (s *InvoiceService) CreateInvoice(req CreateInvoiceRequest) (*Invoice, error) {
	// Parse dates
	issueDate, err := time.Parse("2006-01-02", req.IssueDate)
	if err != nil {
		return nil, errors.New("invalid issue_date format, expected YYYY-MM-DD")
	}

	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		return nil, errors.New("invalid due_date format, expected YYYY-MM-DD")
	}

	invoice := Invoice{
		OrganizationId: req.OrganizationId,
		PaymentId:      req.PaymentId,
		InvoiceNumber:  req.InvoiceNumber,
		CustomerId:     req.CustomerId,
		IssueDate:      issueDate,
		DueDate:        dueDate,
		Amount:         req.Amount,
		TxtAmount:      req.TxtAmount,
	}

	// Save to DB (repository)
	saved, err := s.repo.CreateInvoice(invoice)
	if err != nil {
		return nil, err
	}

	return saved, nil
}

// UpdateInvoice updates invoice fields
func (s *InvoiceService) UpdateInvoice(id string, req UpdateInvoiceRequest) (*Invoice, error) {
	// Load existing
	inv, err := s.repo.GetInvoiceById(id)
	if err != nil {
		return nil, err
	}

	// Patch-style updates
	if req.InvoiceNumber != nil {
		inv.InvoiceNumber = *req.InvoiceNumber
	}
	if req.PaymentId != nil {
		inv.PaymentId = req.PaymentId
	}
	if req.CustomerId != nil {
		inv.CustomerId = *req.CustomerId
	}
	if req.IssueDate != nil {
		parsed, err := time.Parse("2006-01-02", *req.IssueDate)
		if err != nil {
			return nil, errors.New("invalid issue_date format")
		}
		inv.IssueDate = parsed
	}
	if req.DueDate != nil {
		parsed, err := time.Parse("2006-01-02", *req.DueDate)
		if err != nil {
			return nil, errors.New("invalid due_date format")
		}
		inv.DueDate = parsed
	}
	if req.Amount != nil {
		inv.Amount = *req.Amount
	}
	if req.TxtAmount != nil {
		inv.TxtAmount = *req.TxtAmount
	}

	// Update DB
	return s.repo.UpdateInvoice(inv)
}

// DeleteInvoice deletes by ID
func (s *InvoiceService) DeleteInvoice(id string) error {
	return s.repo.DeleteInvoice(id)
}
