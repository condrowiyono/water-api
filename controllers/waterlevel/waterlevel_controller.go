package waterlevel

import (
	"errors"
	"fmt"
	"mini-bank/models"
	"mini-bank/repository"
	"mini-bank/utils"
	"mini-bank/utils/excel"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {
	var waterlevel []*models.WaterLevelObservation
	var filter WaterLevelFilterDTO
	ctx.BindQuery(&filter)

	pagination := utils.Pagination{Page: filter.Page, Limit: filter.Limit}
	total, err := repository.GetWithFilter(&waterlevel, &filter, pagination)

	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	utils.ResponseListSuccess(ctx, waterlevel, utils.Meta{Page: pagination.Page, Limit: pagination.Limit, Total: total})
}

func Create(ctx *gin.Context) {
	var userID = ctx.MustGet("Id").(float64)
	var waterlevel models.WaterLevelObservation

	err := ctx.BindJSON(&waterlevel)
	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	waterlevel.UserID = uint(userID)

	err = repository.Create(&waterlevel)
	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	utils.ResponseCreated(ctx, waterlevel)
}

func GetByID(ctx *gin.Context) {
	var waterlevel models.WaterLevelObservation

	id := ctx.Param("id")
	err := repository.GetByID(&waterlevel, id)

	if err != nil {
		utils.ResponseNotFound(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, waterlevel)

}

func Update(ctx *gin.Context) {
	var waterlevel models.WaterLevelObservation
	id := ctx.Param("id")
	err := repository.GetByID(&waterlevel, id)
	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}
	err = ctx.BindJSON(&waterlevel)
	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	err = repository.Update(&waterlevel)
	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, waterlevel)
}

func Delete(ctx *gin.Context) {
	var waterlevel models.WaterLevelObservation

	id := ctx.Param("id")
	err := repository.GetByID(&waterlevel, id)
	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}
	err = repository.Delete(&waterlevel)

	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, nil)
}

func GetToday(ctx *gin.Context) {
	var waterlevel *[]models.WaterLevelObservation
	riverID := ctx.Param("river")

	err := repository.FindToday(&waterlevel, riverID)

	if err != nil {
		utils.ResponseSuccess(ctx, nil)
		return
	}

	utils.ResponseSuccess(ctx, waterlevel)
}

func ExportByID(ctx *gin.Context) {
	var waterlevel []models.WaterLevelObservation
	var filter WaterLevelExportDTO
	riverID := ctx.Param("river")
	ctx.BindQuery(&filter)

	if filter.Year == "" {
		utils.ResponseBadRequest(ctx, errors.New("year is required"))
		return
	}

	err := repository.ExportObservationByYear(&waterlevel, riverID, filter.Year)

	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	file, err := excel.CreateFile()
	row := 1

	header := []string{
		"Date",
		"Data",
		"Keterangan",
		"Kejadian",
		"Created At",
		"Updated At",
		"Created By",
	}

	excel.SetRow(file, header, row, "")
	row += 1

	for _, d := range waterlevel {

		data := []string{
			d.Date.Format("2006-01-02"),
			fmt.Sprintf("%v", d.Data),
			d.Descrption,
			d.Event,
			d.CreatedAt.Format("2006-01-02 15:04:05"),
			d.UpdatedAt.Format("2006-01-02 15:04:05"),
			d.User.Email,
		}
		excel.SetRow(file, data, row, "")
		row += 1
	}

	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	ctx.Header("Content-Disposition", "attachment; filename=users.xlsx")
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Transfer-Encoding", "binary")

	file.Write(ctx.Writer)
}
