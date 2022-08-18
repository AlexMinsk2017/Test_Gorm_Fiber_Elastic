package repo

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"log"
	"time"
)

type IOrderRepository interface {
	Create(ctx context.Context, dbm *models.Order) (*models.Order, error)
	GetByID(ctx context.Context, id uint) (*models.Order, error)
	DeleteMark(ctx context.Context, id uint) error
	Update(ctx context.Context, dbm *models.Order) (*models.Order, error)
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
func (rep *OrderRepository) DeleteMark(ctx context.Context, id uint) error {
	data, err := rep.GetByID(ctx, id)
	if err != nil {
		return err
	}
	deletedMark := g.IsEmpty(data.DeletedAt)
	if deletedMark == false {
		deletedMark = data.DeletedAt.Valid
	} else {
		deletedMark = false
	}

	data.DeletedAt.Time = time.Now()
	data.DeletedAt.Valid = !deletedMark

	tx := rep.db.WithContext(ctx).Begin()
	err = tx.WithContext(ctx).Unscoped().Save(&data).Error
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	tx.Commit()
	return nil
}
func (rep *OrderRepository) Update(ctx context.Context, dbm *models.Order) (*models.Order, error) {
	tx := rep.db.Begin()
	err := tx.WithContext(ctx).Save(dbm).Error
	//err := rep.db.WithContext(ctx).Create(dbm).Error
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return nil, err
	}
	tx.Commit()
	return dbm, nil
}
