package rainfall

import (
	"errors"
	"fmt"
	"mini-bank/models"
	"mini-bank/repository"
	"mini-bank/utils"
	"mini-bank/utils/excel"
	"time"

	"github.com/gin-gonic/gin"
)

const DEFAULT_TIME_ZONE = "Asia/Singapore"

func GetAll(ctx *gin.Context) {
	var rainfall []*models.RainfallObservation
	var filter RainfallFilterDTO
	ctx.BindQuery(&filter)

	pagination := utils.Pagination{Page: filter.Page, Limit: filter.Limit}
	total, err := repository.GetWithFilter(&rainfall, &filter, pagination)

	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	utils.ResponseListSuccess(ctx, rainfall, utils.Meta{Page: pagination.Page, Limit: pagination.Limit, Total: total})
}

func Create(ctx *gin.Context) {
	var rainfall models.RainfallObservation
	var userID = ctx.MustGet("Id").(float64)

	err := ctx.BindJSON(&rainfall)

	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	rainfall.UserID = uint(userID)

	err = repository.Create(&rainfall)
	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	utils.ResponseCreated(ctx, rainfall)
}

func GetByID(ctx *gin.Context) {
	var rainfall models.RainfallObservation

	id := ctx.Param("id")
	err := repository.GetByID(&rainfall, id)

	if err != nil {
		utils.ResponseNotFound(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, rainfall)
}

func Update(ctx *gin.Context) {
	var rainfall models.RainfallObservation

	id := ctx.Param("id")
	err := repository.GetByID(&rainfall, id)

	if err != nil {
		utils.ResponseNotFound(ctx, err)
		return
	}

	err = ctx.BindJSON(&rainfall)
	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	err = repository.Update(&rainfall)
	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, rainfall)
}

func Delete(ctx *gin.Context) {
	var rainfall models.RainfallObservation

	id := ctx.Param("id")
	err := repository.GetByID(&rainfall, id)

	if err != nil {
		utils.ResponseNotFound(ctx, err)
		return
	}

	err = repository.Delete(&rainfall)
	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, rainfall)
}

func GetToday(ctx *gin.Context) {
	var rainfall models.RainfallObservation
	riverID := ctx.Param("river")

	err := repository.GetToday(&rainfall, riverID)

	if err != nil {
		utils.ResponseSuccess(ctx, nil)
		return
	}

	utils.ResponseSuccess(ctx, rainfall)
}

func ExportByID(ctx *gin.Context) {
	var rainfall []models.RainfallObservation
	var filter ExportDTO
	riverID := ctx.Param("river")
	ctx.BindQuery(&filter)

	if filter.Year == "" {
		utils.ResponseBadRequest(ctx, errors.New("year is required"))
		return
	}

	err := repository.ExportObservationByYear(&rainfall, riverID, filter.Year)

	if err != nil {
		utils.ResponseBadRequest(ctx, err)
		return
	}
	location, _ := time.LoadLocation(DEFAULT_TIME_ZONE)

	file, err := excel.CreateFile()
	row := 1

	header := []string{
		"Date",
		"Data",
		"Durasi",
		"Keterangan",
		"Kejadian",
		"Created At",
		"Updated At",
		"Created By",
	}

	excel.SetRow(file, header, row, "")
	row += 1

	for _, d := range rainfall {

		data := []string{
			d.Date.In(location).Format("2006-01-02"),
			fmt.Sprintf("%v", d.Data),
			fmt.Sprintf("%v", d.Duration),
			d.Descrption,
			d.Event,
			d.CreatedAt.In(location).Format("2006-01-02 15:04:05"),
			d.UpdatedAt.In(location).Format("2006-01-02 15:04:05"),
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
