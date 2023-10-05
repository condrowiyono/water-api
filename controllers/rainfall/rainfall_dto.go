package rainfall

import (
	"mini-bank/utils"
)

type RainfallFilterDTO struct {
	utils.Pagination
	Name string `gorm:"name" form:"name"`
}

type ExportDTO struct {
	Year string `form:"year"`
}
