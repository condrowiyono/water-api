package climate

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
	var climate []*models.ClimateObservation
	var filter ClimateFilterDTO
	ctx.BindQuery(&filter)

	pagination := utils.Pagination{Page: filter.Page, Limit: filter.Limit}
	total, err := repository.GetWithFilter(&climate, &filter, pagination)

	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	utils.ResponseListSuccess(ctx, climate, utils.Meta{Page: pagination.Page, Limit: pagination.Limit, Total: total})
}

func Create(ctx *gin.Context) {
	var climate models.ClimateObservation
	var userID = ctx.MustGet("Id").(float64)

	err := ctx.BindJSON(&climate)
	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}
	climate.UserID = uint(userID)

	err = repository.Create(&climate)
	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	utils.ResponseCreated(ctx, climate)
}

func GetByID(ctx *gin.Context) {
	var climate models.ClimateObservation

	id := ctx.Param("id")
	err := repository.GetByID(&climate, id)

	if err != nil {
		utils.ResponseNotFound(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, climate)

}

func Update(ctx *gin.Context) {
	var climate models.ClimateObservation
	id := ctx.Param("id")

	err := repository.GetByID(&climate, id)

	if err != nil {
		utils.ResponseNotFound(ctx, err)
		return
	}

	err = ctx.BindJSON(&climate)
	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	err = repository.Update(&climate)
	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, climate)
}

func Delete(ctx *gin.Context) {
	var climate models.ClimateObservation

	id := ctx.Param("id")
	err := repository.GetByID(&climate, id)

	if err != nil {
		utils.ResponseNotFound(ctx, err)
		return
	}

	err = repository.Delete(&climate)
	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, climate)
}

func GetToday(ctx *gin.Context) {
	var climate models.ClimateObservation
	riverID := ctx.Param("river")

	err := repository.GetToday(&climate, riverID)

	if err != nil {
		utils.ResponseSuccess(ctx, nil)
		return
	}

	utils.ResponseSuccess(ctx, climate)
}

func ExportByID(ctx *gin.Context) {
	var climate []models.ClimateObservation
	var filter ExportDTO
	riverID := ctx.Param("river")
	ctx.BindQuery(&filter)

	if filter.Year == "" {
		utils.ResponseBadRequest(ctx, errors.New("year is required"))
		return
	}

	err := repository.ExportObservationByYear(&climate, riverID, filter.Year)

	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	file, err := excel.CreateFile()
	row := 1

	header := []string{
		"Date",
		"Min Temperature",
		"Max Temperature",
		"Created At",
		"Updated At",
		"Created By",
	}

	excel.SetRow(file, header, row, "")
	row += 1

	for _, d := range climate {

		data := []string{
			d.Date.Format("2006-01-02"),
			fmt.Sprintf("%v", d.MinTemperature),
			fmt.Sprintf("%v", d.MaxTemperature),
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
