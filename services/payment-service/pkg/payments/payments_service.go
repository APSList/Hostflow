package payments

// PaymentsService handles business logic for payments
type PaymentsService struct {
	repo *PaymentsRepository
}

// NewPaymentsService creates a new PaymentsService
func GetPaymentsService(repo *PaymentsRepository) *PaymentsService {
	return &PaymentsService{
		repo: repo,
	}
}

// GetPayments returns all payments
func (s *PaymentsService) GetPayments() ([]Payment, error) {
	return s.repo.GetPayments()
}
