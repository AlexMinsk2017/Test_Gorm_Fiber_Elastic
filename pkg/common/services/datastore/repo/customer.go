package repo

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"context"
	"gorm.io/gorm"
	"log"
)

type ICustomerRepository interface {
	Create(ctx context.Context, dbm *models.Customer) (*models.Customer, error)
	GetByID(ctx context.Context, id uint) (*models.Customer, error)
}
type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomer(db *gorm.DB) ICustomerRepository {
	return &CustomerRepository{db: db}
}

func (rep *CustomerRepository) Create(ctx context.Context, dbm *models.Customer) (*models.Customer, error) {
	err := rep.db.WithContext(ctx).Create(dbm).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dbm, nil
}
func (rep *CustomerRepository) GetByID(ctx context.Context, id uint) (*models.Customer, error) {
	data := models.Customer{}
	err := rep.db.WithContext(ctx).Unscoped().First(&data, "id = ?", id).Error
	//err := rep.db.WithContext(ctx).
	//	Unscoped().
	//	Model(models.Customer{}).
	//	Table("customers").
	//	Select("*").
	//	Where("customers.id = ?", id).
	//	First(&customer).Error
	return &data, err
}
