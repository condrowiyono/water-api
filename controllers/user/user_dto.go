package user

import "mini-bank/utils"

type UserFilerDTO struct {
	utils.Pagination
	Name      string `gorm:"name" form:"name"`
	Email     string `gorm:"email" form:"email"`
	StartDate int64  `form:"start_date"`
	EndDate   int64  `form:"end_date"`
}

type AttachRoleToUser struct {
	Role []string `json:"roles" binding:"required"`
}
