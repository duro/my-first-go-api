package models

import "github.com/jinzhu/gorm"

// DoMigrate runs the AutoMigrations for GORM
func DoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Author{}, &Publisher{}, &Book{})
}
