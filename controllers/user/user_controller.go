package user

import (
	"fmt"
	"mini-bank/helpers"
	"mini-bank/models"
	"mini-bank/repository"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {
	var user []*models.User
	var filter models.User

	name := ctx.Query("name")
	email := ctx.Query("email")
	facultyEmail := ctx.Query("faculty_email")
	filter = models.User{Name: name, Email: email, FacultyEmail: facultyEmail}
	pagination := helpers.GetPagination(ctx)

	total, err := repository.GetWithFilter(&user, &filter, pagination)

	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	helpers.ResponseListSuccess(ctx, user, helpers.Meta{Page: pagination.Page, Limit: pagination.Limit, Total: total})
}

func Create(ctx *gin.Context) {
	var user models.User

	err := ctx.BindJSON(&user)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	user.Password = helpers.HashPassword(user.Password)

	err = repository.Create(&user)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	helpers.ResponseCreated(ctx, user)
}

func GetByID(ctx *gin.Context) {
	var user models.User

	id := ctx.Param("id")
	err := repository.GetByID(&user, id)

	fmt.Println("err", err)
	if err != nil {
		helpers.ResponseNotFound(ctx, err)
		return
	}

	helpers.ResponseSuccess(ctx, user)
}

func Update(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("id")
	err := repository.GetByID(&user, id)

	if err != nil {
		helpers.ResponseNotFound(ctx, err)
		return
	}

	err = ctx.BindJSON(&user)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	err = repository.Update(&user)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	helpers.ResponseSuccess(ctx, user)
}

func Delete(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("id")
	err := repository.GetByID(&user, id)

	if err != nil {
		helpers.ResponseNotFound(ctx, err)
		return
	}

	err = repository.Delete(&user)
	if err != nil {
		helpers.ResponseBadRequest(ctx, err)
		return
	}

	helpers.ResponseSuccess(ctx, user)
}
