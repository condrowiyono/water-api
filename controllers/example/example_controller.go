package example

import (
	"mini-bank/helpers"
	"mini-bank/models"
	"mini-bank/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetData(ctx *gin.Context) {
	var example []*models.Example
	repository.Get(&example)

	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    example,
	})
}
