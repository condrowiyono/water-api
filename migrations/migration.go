package migrations

import (
	"mini-bank/infra/database"
	"mini-bank/models"
)

// Migrate Add list of model add for migrations
// TODO later separate migration each models
func Migrate() {
	var migrationModels = []interface{}{
		&models.Example{},
		&models.Permission{},
		&models.Role{},
		&models.User{},
	}

	err := database.DB.AutoMigrate(migrationModels...)

	if err != nil {
		return
	}
}
