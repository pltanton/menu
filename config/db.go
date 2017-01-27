package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"kaliwe.ru/menu/db"
	"kaliwe.ru/menu/models"
)

var instance *gorm.DB

func initDB() {
	instance, err := gorm.Open("sqlite3", "db/menu.db")
	if err != nil {
		panic(err)
	}
	instance.AutoMigrate(
		&models.Ingredient{},
		&models.Recipe{},
	)
	db.SetInstance(instance)
}
