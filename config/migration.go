package config

import (
	"tourism-monitoring/entities"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Place{})
	db.AutoMigrate(&entities.VisitReport{})
}
