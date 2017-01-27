package db

import (
	"github.com/jinzhu/gorm"
)

var instance *gorm.DB

func GetInstance() *gorm.DB {
	return instance
}

func SetInstance(newInstance *gorm.DB) {
	instance = newInstance
}
