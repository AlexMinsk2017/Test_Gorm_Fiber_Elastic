package repo

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"context"
	"gorm.io/gorm"
	"log"
)

type IOrderRepository interface {
	Create(ctx context.Context, dbm *models.Order) (*models.Order, error)
	GetByID(ctx context.Context, id uint) (*models.Order, error)
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
func (rep *OrderRepository) GetByID(ctx context.Context, id uint) (*models.Order, error) {
	data := models.Order{}
	err := rep.db.WithContext(ctx).Unscoped().First(&data, "id = ?", id).Error
	return &data, err
}
