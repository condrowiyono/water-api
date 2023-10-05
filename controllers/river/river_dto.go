package river

import "mini-bank/utils"

type RiverFilterDTO struct {
	utils.Pagination
	Name string `gorm:"name" form:"name"`
	Type string `gorm:"type" form:"type"`
}
