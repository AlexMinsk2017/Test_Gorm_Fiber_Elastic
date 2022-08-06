package repo

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"context"
	"gorm.io/gorm"
	"log"
)

type ITableOrderRepository interface {
	Create(ctx context.Context, dbm *models.TableOrder) (*models.TableOrder, error)
	GetByID(ctx context.Context, id uint) (*models.TableOrder, error)
}
type TableOrderRepository struct {
	db *gorm.DB
}

func NewTableOrder(db *gorm.DB) ITableOrderRepository {
	return &TableOrderRepository{db: db}
}

func (rep *TableOrderRepository) Create(ctx context.Context, dbm *models.TableOrder) (*models.TableOrder, error) {
	err := rep.db.WithContext(ctx).Create(dbm).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dbm, nil
}
func (rep *TableOrderRepository) GetByID(ctx context.Context, id uint) (*models.TableOrder, error) {
	data := models.TableOrder{}
	err := rep.db.WithContext(ctx).Unscoped().First(&data, "id = ?", id).Error
	return &data, err
}
