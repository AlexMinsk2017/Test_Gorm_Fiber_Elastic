package repo

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"context"
	"gorm.io/gorm"
	"log"
)

type IGoodRepository interface {
	Create(ctx context.Context, dbm *models.Good) (*models.Good, error)
	GetByID(ctx context.Context, id uint) (*models.Good, error)
}
type GoodRepository struct {
	db *gorm.DB
}

func NewGood(db *gorm.DB) IGoodRepository {
	return &GoodRepository{db: db}
}

func (rep *GoodRepository) Create(ctx context.Context, dbm *models.Good) (*models.Good, error) {
	err := rep.db.WithContext(ctx).Create(dbm).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dbm, nil
}
func (rep *GoodRepository) GetByID(ctx context.Context, id uint) (*models.Good, error) {
	data := models.Good{}
	err := rep.db.WithContext(ctx).Unscoped().First(&data, "id = ?", id).Error
	return &data, err
}
