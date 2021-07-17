package payment

// "gitlab.com/cilalabs/my-sales-champion/service/payment/pkg/logger"

// Service implements Servicer interface.
type Service struct {
	Repo   Repository
	Stripe StripeProvider
}

// NewPaymentService is a factory func to create a new payment.Service.
func NewPaymentService(repo Repository, s StripeProvider) *Service {
	var ps Service
	ps.Repo = repo
	ps.Stripe = s
	return &ps
}
