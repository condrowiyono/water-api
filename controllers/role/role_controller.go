package role

import (
	"fmt"
	"mini-bank/helpers"
	"mini-bank/models"
	"mini-bank/repository"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {
	var role []*models.Role
	var filter models.Role

	name := ctx.Query("name")
	filter = models.Role{Name: name}
	pagination := helpers.GetPagination(ctx)
	preload := "Permissions"

	repository.GetWithFilterWithPreload(&role, &filter, pagination, preload)

	helpers.ResponseSuccess(ctx, role)
}

func Create(ctx *gin.Context) {
	var role models.Role

	err := ctx.BindJSON(&role)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	err = repository.Create(&role)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	helpers.ResponseCreated(ctx, role)
}

func GetByID(ctx *gin.Context) {
	var role models.Role

	id := ctx.Param("id")
	err := repository.GetByIDWithPreload(&role, id, "Permissions")

	fmt.Println("err", err)
	if err != nil {
		helpers.ResponseNotFound(ctx, err)
		return
	}

	helpers.ResponseSuccess(ctx, role)
}

func Update(ctx *gin.Context) {
	var role models.Role
	id := ctx.Param("id")
	err := repository.GetByID(&role, id)

	if err != nil {
		helpers.ResponseNotFound(ctx, err)
		return
	}

	err = ctx.BindJSON(&role)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	err = repository.Update(&role)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	helpers.ResponseSuccess(ctx, role)
}

func Delete(ctx *gin.Context) {
	var role models.Role
	id := ctx.Param("id")
	err := repository.GetByID(&role, id)

	if err != nil {
		helpers.ResponseNotFound(ctx, err)
		return
	}

	err = repository.Delete(&role)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	helpers.ResponseSuccess(ctx, role)
}

func AddPermissions(ctx *gin.Context) {
	var permissions AddPermissionsToRole
	id := ctx.Param("id")
	err := ctx.BindJSON(&permissions)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	err = repository.AddPermissionsToRole(id, permissions.Permission)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	helpers.ResponseSuccess(ctx, nil)
}
