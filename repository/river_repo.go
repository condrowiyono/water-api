package repository

import (
	"mini-bank/infra/database"
	"mini-bank/utils"
	"time"
)

func ExportObservationByYear(model interface{}, id, year string) error {
	query := database.DB.Model(model).Where("river_id = ?", id).Where("date LIKE ?", year+"%").Preload("User").Find(model)

	return query.Error
}

func GetToday(model interface{}, id string) error {
	startTime := utils.StartOfDay(time.Now(), "Asia/Makassar").UTC()
	endTime := utils.EndOfDay(time.Now(), "Asia/Makassar").UTC()

	query := database.DB.Model(model).Where("river_id = ?", id).Where("date BETWEEN ? AND ?", startTime, endTime).Preload("User").First(model)

	return query.Error
}

func FindToday(model interface{}, id string) error {
	startTime := utils.StartOfDay(time.Now(), "Asia/Makassar").UTC()
	endTime := utils.EndOfDay(time.Now(), "Asia/Makassar").UTC()

	query := database.DB.Model(model).Where("river_id = ?", id).Where("date BETWEEN ? AND ?", startTime, endTime).Preload("User").Find(model)

	return query.Error
}
