package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	User string `gorm:"size:50"`
	Pass string `gorm:"size:50"`
}
