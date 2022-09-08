package repo

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"context"
	"gorm.io/gorm"
	"log"
)

type IUserRepository interface {
	Create(ctx context.Context, dbm *models.Users) (*models.Users, error)
}
type UserRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (rep *UserRepository) Create(ctx context.Context, dbm *models.Users) (*models.Users, error) {
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
