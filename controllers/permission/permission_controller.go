package permission

import (
	"fmt"
	"mini-bank/helpers"
	"mini-bank/models"
	"mini-bank/repository"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {
	var permission []*models.Permission
	var filter models.Permission

	name := ctx.Query("name")
	filter = models.Permission{Name: name}
	pagination := helpers.GetPagination(ctx)
	repository.GetWithFilter(&permission, &filter, pagination)

	helpers.ResponseSuccess(ctx, permission)
}

func Create(ctx *gin.Context) {
	var permission models.Permission

	err := ctx.BindJSON(&permission)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	err = repository.Create(&permission)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	helpers.ResponseCreated(ctx, permission)
}

func GetByID(ctx *gin.Context) {
	var permission models.Permission

	id := ctx.Param("id")
	err := repository.GetByID(&permission, id)

	fmt.Println("err", err)
	if err != nil {
		helpers.ResponseNotFound(ctx, err)
		return
	}

	helpers.ResponseSuccess(ctx, permission)
}

func Update(ctx *gin.Context) {
	var permission models.Permission
	id := ctx.Param("id")
	err := repository.GetByID(&permission, id)

	if err != nil {
		helpers.ResponseNotFound(ctx, err)
		return
	}

	err = ctx.BindJSON(&permission)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	err = repository.Update(&permission)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	helpers.ResponseSuccess(ctx, permission)
}

func Delete(ctx *gin.Context) {
	var permission models.Permission
	id := ctx.Param("id")
	err := repository.GetByID(&permission, id)

	if err != nil {
		helpers.ResponseNotFound(ctx, err)
		return
	}

	err = repository.Delete(&permission)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	helpers.ResponseSuccess(ctx, permission)
}
