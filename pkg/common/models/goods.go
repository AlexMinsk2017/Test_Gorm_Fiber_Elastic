package models

import "gorm.io/gorm"

type Good struct {
	gorm.Model
	Name    string `gorm:"size:150"`
	Price   float32
	Comment string `gorm:"size:300"`
}
