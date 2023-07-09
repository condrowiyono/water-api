package models

import (
	"time"
)

type Permission struct {
	Id          string     `gorm:"type:VARCHAR(36);primary_key;" json:"id"`
	Name        string     `gorm:"unique" json:"name" binding:"required"`
	Description string     `json:"description" binding:"required"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

// TableName is Database TableName of this model
func (permission *Permission) TableName() string {
	return "permissions"
}
