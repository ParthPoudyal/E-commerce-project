package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ParthPoudyal/E-commerce-project/infrastructure/mapper"
	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) AddUser(user *domain.User) error {
	model := mapper.ToUserModel(user)
	if model.ID == uuid.Nil {
		model.ID = uuid.New()
	}

	result := r.db.Create(model)
	if result.Error != nil {
		return fmt.Errorf("create user: %w", result.Error)
	}

	user.ID = model.ID
	user.CreatedAt = model.CreatedAt
	user.UpdatedAt = model.UpdatedAt

	return nil
}

func (r *PostgresUserRepository) FindUserByID(id uuid.UUID) (*domain.User, error) {
	var model models.UserModel
	result := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("find user by id %s: %w", id, result.Error)
	}

	return mapper.ToUserDomain(&model), nil
}

func (r *PostgresUserRepository) FindUserByEmail(email string) (*domain.User, error) {
	var model models.UserModel
	result := r.db.Where("email = ? AND deleted_at IS NULL", email).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("find user by email %s: %w", email, result.Error)
	}

	return mapper.ToUserDomain(&model), nil
}

func (r *PostgresUserRepository) RemoveUser(id uuid.UUID) error {
	now := time.Now()
	result := r.db.Model(&models.UserModel{}).Where("id = ?", id).Update("deleted_at", &now)
	if result.Error != nil {
		return fmt.Errorf("remove user %s: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

func (r *PostgresUserRepository) UpdateUser(user *domain.User) error {
	model := mapper.ToUserModel(user)
	result := r.db.Model(&models.UserModel{}).Where("id = ? AND deleted_at IS NULL", user.ID).Updates(map[string]interface{}{
		"username":   model.Username,
		"address":    model.Address,
		"phone_no":   model.PhoneNo,
		"email":      model.Email,
		"updated_at": time.Now(),
	})
	if result.Error != nil {
		return fmt.Errorf("update user %s: %w", user.ID, result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

func (r *PostgresUserRepository) SaveAuthCredential(cred *domain.AuthCredential) error {
	model := mapper.ToAuthCredentialModel(cred)
	result := r.db.Save(model)
	if result.Error != nil {
		return fmt.Errorf("save auth credential: %w", result.Error)
	}
	return nil
}

func (r *PostgresUserRepository) FindAuthCredentialByUserID(userID uuid.UUID) (*domain.AuthCredential, error) {
	var model models.AuthCredentialModel
	result := r.db.Where("user_id = ?", userID).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrInvalidCredentials
		}
		return nil, fmt.Errorf("find auth credential for user %s: %w", userID, result.Error)
	}
	return mapper.ToAuthCredentialDomain(&model), nil
}

