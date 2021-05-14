package configs

import (
	"assignment_2/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, reset bool) {
	if reset == true {
		_ = db.Migrator().DropTable(&models.Item{})
		_ = db.Migrator().DropTable(&models.Order{})
	}
	_ = db.Migrator().CreateTable(&models.Order{})
	_ = db.Migrator().CreateTable(&models.Item{})
}