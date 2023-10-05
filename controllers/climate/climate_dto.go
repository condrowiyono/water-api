package climate

import (
	"mini-bank/utils"
)

type ClimateFilterDTO struct {
	utils.Pagination
	Name string `gorm:"name" form:"name"`
}

type ExportDTO struct {
	Year string `form:"year"`
}
