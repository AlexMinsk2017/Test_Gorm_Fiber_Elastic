package web

import "Test_Gorm_Fiber_Elastic/pkg/common/models/dto"

type Good struct {
	Id          uint
	Name        string
	Comment     string
	Price       float32
	DeletedMark bool
}

type Goods []*Good

func (r *Good) ToDTO() *dto.Good {
	if r == nil {
		return nil
	}
	return &dto.Good{
		Id:      r.Id,
		Name:    r.Name,
		Comment: r.Comment,
		Price:   r.Price,
	}
}
func (r *Good) FromDTO(v *dto.Good) *Good {
	if v == nil {
		return nil
	}
	return &Good{
		Id:      v.Id,
		Name:    v.Name,
		Comment: v.Comment,
		Price:   v.Price,
	}
}
