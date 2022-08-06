package repo

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"context"
	"gorm.io/gorm"
	"log"
)

type IOrderRepository interface {
	Create(ctx context.Context, dbm *models.Order) (*models.Order, error)
}
type OrderRepository struct {
	db *gorm.DB
}

func NewOrder(db *gorm.DB) IOrderRepository {
	return &OrderRepository{db: db}
}

func (rep *OrderRepository) Create(ctx context.Context, dbm *models.Order) (*models.Order, error) {
	err := rep.db.WithContext(ctx).Create(dbm).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dbm, nil
}
