package user

import (
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/google/uuid"
)

type UserRepo interface {
	// To Add User in Db
	AddUser(user *domain.User) error

	// to Find User by ID, returns User
	FindUserByID(id uuid.UUID) (*domain.User, error)

	// to Find User by email, returns User
	FindUserByEmail(mail string) (*domain.User, error)

	// to remove User using id, if exists
	RemoveUser(id uuid.UUID) error

	// to Update User data, if exists
	UpdateUser(user *domain.User) error

	// Auth credential methods
	SaveAuthCredential(cred *domain.AuthCredential) error
	FindAuthCredentialByUserID(userID uuid.UUID) (*domain.AuthCredential, error)
}
