package repo

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"log"
	"time"
)

type ICustomerRepository interface {
	Create(ctx context.Context, dbm *models.Customer) (*models.Customer, error)
	GetByID(ctx context.Context, id uint) (*models.Customer, error)
	DeleteMark(ctx context.Context, id uint) error
	Update(ctx context.Context, dbm *models.Customer) (*models.Customer, error)
	WithTx(tx *gorm.DB) ICustomerRepository
	SelectMaxID(ctx context.Context) (uint, error)
}
type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomer(db *gorm.DB) ICustomerRepository {
	return &CustomerRepository{db: db}
}

func (rep *CustomerRepository) WithTx(tx *gorm.DB) ICustomerRepository {
	return NewCustomer(tx)
}

func (rep *CustomerRepository) Create(ctx context.Context, dbm *models.Customer) (*models.Customer, error) {
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
func (rep *CustomerRepository) DeleteMark(ctx context.Context, id uint) error {

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
func (rep *CustomerRepository) Update(ctx context.Context, dbm *models.Customer) (*models.Customer, error) {
	tx := rep.db.Begin()
	tx.WithContext(ctx).Unscoped().First(&dbm)
	err := tx.WithContext(ctx).Unscoped().Save(dbm).Error
	//err := rep.db.WithContext(ctx).Save(dbm).Error
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return nil, err
	}
	tx.Commit()
	return dbm, nil
}
func (rep *CustomerRepository) SelectMaxID(ctx context.Context) (uint, error) {

	//data := models.Customer{}
	//db := rep.db.WithContext(ctx).Unscoped().Exec("SELECT MAX(id) FROM customers", &data)
	//fmt.Print(db)
	var maxID uint
	rep.db.WithContext(ctx).Unscoped().Raw("SELECT MAX(id) FROM customers").Scan(&maxID)

	return maxID, nil
}
