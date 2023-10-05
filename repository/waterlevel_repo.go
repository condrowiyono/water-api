package repository

import (
	"mini-bank/infra/database"
)

func ExportObservationByYear(model interface{}, id, year string) error {
	query := database.DB.Model(model).Where("river_id = ?", id).Where("date LIKE ?", year+"%").Preload("User").Find(model)

	return query.Error
}
