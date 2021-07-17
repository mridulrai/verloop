package storage

import (
	"github.com/mridulrai/verloop/pkg/payment"
)

const collectionPayments string = "payments"

// Payment represents a storage related payment with bson struct tags.
type Payment struct {
}

// Import maps a payment.Payment to StoragePayment.
func (sp *Storage) Import(u payment.Payment) {
}

// Export maps a StoragePayment to payment.Payment.
func (sp *Storage) Export() *payment.Payment {
	var p payment.Payment

	return &p
}
