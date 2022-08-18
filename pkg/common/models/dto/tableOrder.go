package dto

type TableOrder struct {
	Id          uint
	OrderRefer  uint
	GoodRefer   uint
	Quantity    float32
	Price       float32
	Summa       float32
	DeletedMark bool
}
