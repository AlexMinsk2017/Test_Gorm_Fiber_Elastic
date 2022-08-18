package web

import "Test_Gorm_Fiber_Elastic/pkg/common/models/dto"

type TableOrder struct {
	Id          uint
	OrderRefer  uint
	GoodRefer   uint
	Quantity    float32
	Price       float32
	Summa       float32
	DeletedMark bool
}

type TableOrders []*TableOrder

func (r *TableOrder) ToDTO() *dto.TableOrder {
	if r == nil {
		return nil
	}
	return &dto.TableOrder{
		Id:         r.Id,
		OrderRefer: r.OrderRefer,
		GoodRefer:  r.GoodRefer,
		Quantity:   r.Quantity,
		Price:      r.Price,
		Summa:      r.Summa,
	}
}
func (r *TableOrder) FromDTO(v *dto.TableOrder) *TableOrder {
	if v == nil {
		return nil
	}
	return &TableOrder{
		Id:         v.Id,
		OrderRefer: v.OrderRefer,
		GoodRefer:  v.GoodRefer,
		Quantity:   v.Quantity,
		Price:      v.Price,
		Summa:      v.Summa,
	}
}
