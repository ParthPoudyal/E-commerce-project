package mapper

import (
	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
)

func ToPaymentDomain(model *models.PaymentTransactionModel) *domain.PaymentTransaction {
	if model == nil {
		return nil
	}

	return &domain.PaymentTransaction{
		ID:            model.ID,
		PurchaseID:    model.PurchaseID,
		ItemID:        model.ItemID,
		UserID:        model.UserID,
		Amount:        model.Amount,
		Provider:      model.Provider,
		ProviderTxnID: model.ProviderTxnID,
		Status:        domain.PaymentStatus(model.Status),
		PaymentURL:    model.PaymentURL,
		CreatedAt:     model.CreatedAt,
		CompletedAt:   model.CompletedAt,
	}
}

func ToPaymentModel(entity *domain.PaymentTransaction) *models.PaymentTransactionModel {
	if entity == nil {
		return nil
	}

	return &models.PaymentTransactionModel{
		ID:            entity.ID,
		PurchaseID:    entity.PurchaseID,
		ItemID:        entity.ItemID,
		UserID:        entity.UserID,
		Amount:        entity.Amount,
		Provider:      entity.Provider,
		ProviderTxnID: entity.ProviderTxnID,
		Status:        string(entity.Status),
		PaymentURL:    entity.PaymentURL,
		CreatedAt:     entity.CreatedAt,
		CompletedAt:   entity.CompletedAt,
	}
}

func ToPaymentDomainSlice(models []*models.PaymentTransactionModel) []*domain.PaymentTransaction {
	if models == nil {
		return nil
	}

	res := make([]*domain.PaymentTransaction, len(models))
	for i, m := range models {
		res[i] = ToPaymentDomain(m)
	}
	return res
}

