package models

import (
	"time"
)

type User struct {
	Id                   string     `gorm:"type:VARCHAR(36);primary_key;" json:"id"`
	Email                string     `gorm:"unique" json:"email" binding:"required"`
	Password             string     `json:"password,omitempty" binding:"required"`
	FacultyEmail         string     `gorm:"unique" json:"faculty_email" binding:"required"`
	Name                 string     `json:"name" `
	VerificationToken    string     `json:"verification_token,omitempty"`
	VerificationTokenExp *time.Time `json:"verification_token_exp,omitempty"`
	IsVerified           *bool      `json:"is_verified"`
	PasswordToken        string     `json:"password_token,omitempty"`
	PasswordTokenExp     *time.Time `json:"password_token_exp,omitempty"`
	CreatedAt            *time.Time `json:"created_at"`
	UpdatedAt            *time.Time `json:"updated_at"`
	DeletedAt            *time.Time `json:"deleted_at,omitempty"`
	Roles                []Role     `gorm:"many2many:user_roles;" json:"roles,omitempty"`
}
