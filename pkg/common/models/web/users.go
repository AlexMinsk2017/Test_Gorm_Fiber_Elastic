package web

import "Test_Gorm_Fiber_Elastic/pkg/common/models/dto"

type User struct {
	Id   uint
	User string
	Pass string
}

type Users []*User

func (r *User) ToDTO() *dto.User {
	if r == nil {
		return nil
	}
	return &dto.User{
		Id:   r.Id,
		User: r.User,
		Pass: r.Pass,
	}
}
func (r *User) FromDTO(v *dto.User) *User {
	if v == nil {
		return nil
	}
	return &User{
		Id:   v.Id,
		User: v.User,
		Pass: v.Pass,
	}
}
