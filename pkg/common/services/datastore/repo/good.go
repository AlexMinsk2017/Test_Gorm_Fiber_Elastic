package repo

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"log"
	"time"
)

type IGoodRepository interface {
	Create(ctx context.Context, dbm *models.Good) (*models.Good, error)
	GetByID(ctx context.Context, id uint) (*models.Good, error)
	DeleteMark(ctx context.Context, id uint) error
	Update(ctx context.Context, dbm *models.Good) (*models.Good, error)
}
type GoodRepository struct {
	db *gorm.DB
}

func NewGood(db *gorm.DB) IGoodRepository {
	return &GoodRepository{db: db}
}

func (rep *GoodRepository) Create(ctx context.Context, dbm *models.Good) (*models.Good, error) {
	tx := rep.db.Begin()
	err := tx.WithContext(ctx).Create(dbm).Error
	//err := rep.db.WithContext(ctx).Create(dbm).Error
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return nil, err
	}
	tx.Commit()
	return dbm, nil
}
func (rep *GoodRepository) GetByID(ctx context.Context, id uint) (*models.Good, error) {
	data := models.Good{}
	err := rep.db.WithContext(ctx).Unscoped().First(&data, "id = ?", id).Error
	return &data, err
}
func (rep *GoodRepository) DeleteMark(ctx context.Context, id uint) error {

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
func (rep *GoodRepository) Update(ctx context.Context, dbm *models.Good) (*models.Good, error) {
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
