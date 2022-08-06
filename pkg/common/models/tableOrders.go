package models

import (
	"gorm.io/gorm"
)

type TableOrder struct {
	gorm.Model
	OrderRefer uint
	Order      Order `gorm:"foreignKey:OrderRefer"`
	GoodRefer  uint
	Good       Good `gorm:"foreignKey:GoodRefer"`
	Quantity   float32
	Price      float32
	Summa      float32
}
