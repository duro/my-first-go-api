package models

import "github.com/jinzhu/gorm"

// Publisher model
type Publisher struct {
	gorm.Model

	Name          string
	StreetAddress string
	City          string
	State         string
	PhoneNumber   string

	Books []Book `gorm:"many2many:publisher_books;"`
}
