package payment

import (
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/google/uuid"
)

type PaymentRepo interface {
	CreatePayment(tx *domain.PaymentTransaction) error
	FindByID(id uuid.UUID) (*domain.PaymentTransaction, error)
	FindByProviderTxnID(pidx string) (*domain.PaymentTransaction, error)
	UpdatePayment(tx *domain.PaymentTransaction) error
	ListPaymentsByUser(userID uuid.UUID) ([]*domain.PaymentTransaction, error)
}

