package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Config() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./config/gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db

}
