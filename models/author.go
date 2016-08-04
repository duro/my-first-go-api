package models

import "github.com/jinzhu/gorm"

// Author model
type Author struct {
	gorm.Model

	FirstName string
	LastName  string
	Books     []Book
}
