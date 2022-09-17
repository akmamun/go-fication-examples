package migrations

import (
	"go-fication-examples/infra/database"
	"go-fication-examples/models"
)

func Migrate(db *database.DB) {
	var migrationModels = []interface{}{&models.Example{}}
	err := db.Database.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
