package repository

import (
	"mini-bank/helpers"
	"mini-bank/infra/database"
	"mini-bank/models"
)

func GetAllUser(model models.User, filter interface{}, pagination helpers.Pagination) (int64, error) {
	var count int64

	query := database.DB

	if filter != nil {
		query = query.Where(filter)
	}

	query.Model(model).Count(&count)

	return count, query.Scopes(helpers.WithPagination(pagination)).Order("created_at DESC").Find(model).Error
}
