package repo

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"context"
	"gorm.io/gorm"
	"log"
)

type ICustomerRepository interface {
	Create(ctx context.Context, dbm *models.Customer) (*models.Customer, error)
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
