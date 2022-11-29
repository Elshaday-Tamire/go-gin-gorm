package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string
	Description string
	Base64      string
}

type Order struct {
	gorm.Model
	OrderName string
}
