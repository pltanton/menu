package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"kaliwe.ru/menu/models"
)

var instance *gorm.DB

func GetInstance() *gorm.DB {
	if instance == nil {
		db, err := gorm.Open("sqlite3", "db/menu.db")
		if err != nil {
			panic(err)
		}
		db.AutoMigrate(
			&models.Ingredient{},
			&models.Recipe{},
		)
		instance = db
	}
	return instance
}
