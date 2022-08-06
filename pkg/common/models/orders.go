package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Number        string `gorm:"size:15"`
	Date          time.Time
	CustomerRefer uint
	Customer      Customer `gorm:"foreignKey:CustomerRefer"`
	Comment       string   `gorm:"size:300"`
}
