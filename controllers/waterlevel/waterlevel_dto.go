package waterlevel

import "mini-bank/utils"

type WaterLevelFilterDTO struct {
	utils.Pagination
	Name string `gorm:"name" form:"name"`
}

type WaterLevelExportDTO struct {
	Year string `form:"year"`
}
