package utils

import (
	"reflect"

	"github.com/google/uuid"
)

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
