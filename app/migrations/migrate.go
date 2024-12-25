package migrations

import (
	"go-elastic/config"
	"go-elastic/models"
)

func RunMigrations() {
	err := config.DB.AutoMigrate(&models.Product{})
	if err != nil {
		panic("Failed to run migrations!")
	}
}
