package web

import "Test_Gorm_Fiber_Elastic/pkg/common/models/dto"

type Customer struct {
	Id      uint
	Name    string
	Comment string
}

type Customers []*Customer

func (r *Customer) ToDTO() *dto.Customer {
	if r == nil {
		return nil
	}
	return &dto.Customer{
		Id:      r.Id,
		Name:    r.Name,
		Comment: r.Comment,
	}
}
func (r *Customer) FromDTO(v *dto.Customer) *Customer {
	if v == nil {
		return nil
	}
	return &Customer{
		Id:      v.Id,
		Name:    v.Name,
		Comment: v.Comment,
	}
}
