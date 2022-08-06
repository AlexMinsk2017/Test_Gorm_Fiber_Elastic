package web

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models/dto"
	"time"
)

type Order struct {
	Id            uint
	Number        string
	Date          *time.Time
	CustomerRefer uint
	Comment       string
}

type Orders []*Order

func (r *Order) ToDTO() *dto.Order {
	if r == nil {
		return nil
	}
	return &dto.Order{
		Id:            r.Id,
		Number:        r.Number,
		Date:          r.Date,
		CustomerRefer: r.CustomerRefer,
		Comment:       r.Comment,
	}
}
func (r *Order) FromDTO(v *dto.Order) *Order {
	if v == nil {
		return nil
	}
	return &Order{
		Id:            v.Id,
		Number:        v.Number,
		Date:          r.Date,
		CustomerRefer: r.CustomerRefer,
		Comment:       r.Comment,
	}
}
