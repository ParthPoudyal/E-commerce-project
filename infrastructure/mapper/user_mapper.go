package mapper

import (
	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
)

func ToUserDomain(model *models.UserModel) *domain.User {
	if model == nil {
		return nil
	}

	return &domain.User{
		ID:       model.ID,
		Username: model.Username,
		Address:  model.Address,
		PhoneNo:  model.PhoneNo,
		Email:    model.Email,
		Base: domain.Base{
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
			DeletedAt: model.DeletedAt,
		},
	}
}

func ToUserModel(entity *domain.User) *models.UserModel {
	if entity == nil {
		return nil
	}

	return &models.UserModel{
		ID:        entity.ID,
		Username:  entity.Username,
		Address:   entity.Address,
		PhoneNo:   entity.PhoneNo,
		Email:     entity.Email,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}

func ToAuthCredentialDomain(model *models.AuthCredentialModel) *domain.AuthCredential {
	if model == nil {
		return nil
	}

	return &domain.AuthCredential{
		UserID:       model.UserID,
		PasswordHash: model.PasswordHash,
		Provider:     model.Provider,
		ProviderID:   model.ProviderID,
	}
}

func ToAuthCredentialModel(entity *domain.AuthCredential) *models.AuthCredentialModel {
	if entity == nil {
		return nil
	}

	return &models.AuthCredentialModel{
		UserID:       entity.UserID,
		PasswordHash: entity.PasswordHash,
		Provider:     entity.Provider,
		ProviderID:   entity.ProviderID,
	}
}
