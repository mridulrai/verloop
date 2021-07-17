package payment

// Payment defines the properties of a payment.
type Payment struct {
}

// Repository defines the methods a storage implementation must support.
type Repository interface {
}

// Servicer defines the methods a payment service implementation must support.
type Servicer interface {
}

// StripeProvider defines the methods a storage implementation must support.
type StripeProvider interface {
}
