package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name    string `gorm:"size:150"`
	Comment string `gorm:"size:300"`
}
