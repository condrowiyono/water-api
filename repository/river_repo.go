package repository

import (
	"mini-bank/infra/database"
	"mini-bank/utils"
	"time"
)

const DEFAULT_TIME_ZONE = "Asia/Singapore"

func ExportObservationByYear(model interface{}, id, year string) error {
	query := database.DB.Model(model).Where("river_id = ?", id).Where("date LIKE ?", year+"%").Preload("User").Find(model)

	return query.Error
}

func GetToday(model interface{}, id string) error {
	startTime := utils.StartOfDay(time.Now(), DEFAULT_TIME_ZONE).UTC()
	endTime := utils.EndOfDay(time.Now(), DEFAULT_TIME_ZONE).UTC()

	query := database.DB.Model(model).Where("river_id = ?", id).Where("date BETWEEN ? AND ?", startTime, endTime).Preload("User").First(model)

	return query.Error
}

func FindToday(model interface{}, id string) error {
	startTime := utils.StartOfDay(time.Now(), DEFAULT_TIME_ZONE).UTC()
	endTime := utils.EndOfDay(time.Now(), DEFAULT_TIME_ZONE).UTC()

	query := database.DB.Model(model).Where("river_id = ?", id).Where("date BETWEEN ? AND ?", startTime, endTime).Preload("User").Find(model)

	return query.Error
}
