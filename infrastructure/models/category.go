package models

import (
	"github.com/google/uuid"
)

type CategoryModel struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string    `gorm:"type:varchar(100);not null;uniqueIndex"`
	Slug        string    `gorm:"type:varchar(100);not null;uniqueIndex"`
	Description string    `gorm:"type:text"`
	ImageURL    string    `gorm:"type:text"`
}

func (CategoryModel) TableName() string {
	return "categories"
}

