package v1

import "github.com/mridulrai/verloop/pkg/verloop"

// PaymentRequest represents Payment request data.
type PaymentRequest struct {
	Data PaymentRequestBody `json:"data"`
}

// PaymentRequestBody represents a PaymentRequest profile.
type PaymentRequestBody struct {
	Type string `json:"type"` // Type of resource
}

func MapPaymentRequestToPaymentRepo(pr PaymentRequest) payment.Payment {
	var p payment.Payment

	return p
}
