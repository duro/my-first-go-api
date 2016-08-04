package models

import "github.com/jinzhu/gorm"

// Book model
type Book struct {
	gorm.Model

	Title       string
	Description string
	Author      Author
}
