package v1

import (
	"github.com/mridulrai/verloop/pkg/payment"
)

// PaymentResponse represents Payment response data.
type PaymentResponse struct {
	Data PaymentResponseBody `json:"data"`
}

// PaymentResponseBody represents a PaymentResponse profile.
type PaymentResponseBody struct {
	Type string `json:"type"` // Type of resource
}

func MapPaymentRepoToPaymentResponse(p payment.Payment) PaymentResponse {
	var pr PaymentResponse

	return pr
}
