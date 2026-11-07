package repository

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ParthPoudyal/E-commerce-project/infrastructure/mapper"
	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
)

type PostgresPaymentRepository struct {
	db *gorm.DB
}

func NewPostgresPaymentRepository(db *gorm.DB) *PostgresPaymentRepository {
	return &PostgresPaymentRepository{
		db: db,
	}
}

func (r *PostgresPaymentRepository) CreatePayment(tx *domain.PaymentTransaction) error {
	model := mapper.ToPaymentModel(tx)
	if model.ID == uuid.Nil {
		model.ID = uuid.New()
	}

	result := r.db.Create(model)
	if result.Error != nil {
		return fmt.Errorf("create payment transaction: %w", result.Error)
	}

	tx.ID = model.ID
	tx.CreatedAt = model.CreatedAt
	return nil
}

func (r *PostgresPaymentRepository) FindByID(id uuid.UUID) (*domain.PaymentTransaction, error) {
	var model models.PaymentTransactionModel
	result := r.db.Where("id = ?", id).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrPaymentNotFound
		}
		return nil, fmt.Errorf("find payment by id %s: %w", id, result.Error)
	}
	return mapper.ToPaymentDomain(&model), nil
}

func (r *PostgresPaymentRepository) FindByProviderTxnID(pidx string) (*domain.PaymentTransaction, error) {
	var model models.PaymentTransactionModel
	result := r.db.Where("provider_txn_id = ?", pidx).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrPaymentNotFound
		}
		return nil, fmt.Errorf("find payment by pidx %s: %w", pidx, result.Error)
	}
	return mapper.ToPaymentDomain(&model), nil
}

func (r *PostgresPaymentRepository) UpdatePayment(tx *domain.PaymentTransaction) error {
	model := mapper.ToPaymentModel(tx)
	result := r.db.Model(&models.PaymentTransactionModel{}).Where("id = ?", tx.ID).Updates(map[string]interface{}{
		"purchase_id":     model.PurchaseID,
		"provider_txn_id": model.ProviderTxnID,
		"status":          model.Status,
		"payment_url":     model.PaymentURL,
		"completed_at":    model.CompletedAt,
	})
	if result.Error != nil {
		return fmt.Errorf("update payment %s: %w", tx.ID, result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrPaymentNotFound
	}
	return nil
}

func (r *PostgresPaymentRepository) ListPaymentsByUser(userID uuid.UUID) ([]*domain.PaymentTransaction, error) {
	var modelList []*models.PaymentTransactionModel
	result := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&modelList)
	if result.Error != nil {
		return nil, fmt.Errorf("list payments by user %s: %w", userID, result.Error)
	}
	return mapper.ToPaymentDomainSlice(modelList), nil
}

