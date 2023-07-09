package helpers

import (
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func GetPagination(ctx *gin.Context) Pagination {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	if page < 0 {
		page = 1
	}

	if limit < 0 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}

	return Pagination{Page: page, Limit: limit}
}

func WithPagination(pagination Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := pagination.Page

		limit := pagination.Limit

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

func AddUUIDToModel(model interface{}) {
	uuidStr := uuid.NewString()

	// Use reflection to set the `Id` field of the model to the generated UUID
	value := reflect.ValueOf(model).Elem()
	idField := value.FieldByName("Id")

	// Check if the `Id` field exists and is assignable
	if idField.IsValid() && idField.CanSet() && idField.Type().String() == "string" {
		idField.SetString(uuidStr)
	}
}
