package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	//user description
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username string     `gorm:"uniqueIndex;not null"`
	Address  string     `gorm:"uniqueIndex;not null"`
	PhoneNo  string
	Email    string

	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`

}

func (UserModel) TableName() string {
	 return "users" 
}

type AuthCredentialModel struct {
    UserID       uuid.UUID `gorm:"type:uuid;primaryKey"`
    PasswordHash *string
    Provider     string  `gorm:"not null;default:'local'"`
    ProviderID   *string `gorm:"index"`
}

func (AuthCredentialModel) TableName() string {
	 return "auth_credentials" 
}