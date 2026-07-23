package entitites

import (
	"github.com/google/uuid"
)

// User : Relational model
type User struct {
	//user description
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Address  string    `json:"address"`
	PhoneNo  string    `json:"phone_no"`
	Email    string    `json:"email"`

	//base
	Base
}

type AuthCredential struct {
	UserID       uuid.UUID
	PasswordHash *string
	Provider     string
	ProviderID   *string
}

// entity attributes and elements
