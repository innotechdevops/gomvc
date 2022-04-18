package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connection(dataSource string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
