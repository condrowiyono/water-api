package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                   uint              `gorm:"primarykey" json:"id"`
	Email                string            `gorm:"unique" json:"email" binding:"required"`
	Password             string            `json:"password,omitempty" binding:"required"`
	Name                 string            `json:"name" `
	VerificationToken    string            `json:"verification_token,omitempty"`
	VerificationTokenExp *time.Time        `json:"verification_token_exp,omitempty"`
	IsVerified           *bool             `json:"is_verified"`
	PasswordToken        string            `json:"password_token,omitempty"`
	PasswordTokenExp     *time.Time        `json:"password_token_exp,omitempty"`
	Roles                []Role            `gorm:"many2many:user_roles;" json:"roles,omitempty"`
	RiverID              uint              `json:"river_id" form:"river_id"`
	RiverType            string            `gorm:"-" json:"river_type" form:"river_type"`
	River                *RiverObservation `json:"river"`
	CreatedAt            time.Time         `json:"created_at"`
	UpdatedAt            time.Time         `json:"updated_at"`
	DeletedAt            gorm.DeletedAt    `gorm:"index" json:"deleted_at"`
}
