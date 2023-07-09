package models

import (
	"time"
)

type Role struct {
	Id          string       `gorm:"type:VARCHAR(36);primary_key;" json:"id"`
	Name        string       `gorm:"unique" json:"name" binding:"required"`
	Description string       `json:"description" binding:"required"`
	CreatedAt   *time.Time   `json:"created_at"`
	UpdatedAt   *time.Time   `json:"updated_at"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}

// TableName is Database TableName of this model
func (role *Role) TableName() string {
	return "roles"
}
