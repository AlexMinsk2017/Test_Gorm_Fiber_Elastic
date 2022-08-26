package orchestrator

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"Test_Gorm_Fiber_Elastic/pkg/common/models/dto"
	"Test_Gorm_Fiber_Elastic/pkg/common/models/web"
	"Test_Gorm_Fiber_Elastic/pkg/common/services/io/elastic/controllers/model"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"log"
)

type ICustomerOrchestrator interface {
	GetByID(ctx context.Context, id uint) (*models.Customer, error)
	Create(ctx context.Context, body *web.Customer) (*models.Customer, error)
	DeleteMark(ctx context.Context, id uint) error
	Update(ctx context.Context, body *web.Customer) (*models.Customer, error)
	UpdateElastic(ctx context.Context) error
	Search(ctx context.Context, value string) (*model.SearchResponse, error)
}
type CustomerOrchestrator struct {
	Engine *Engine
}

func NewCustomerOrchestrator(engine *Engine) ICustomerOrchestrator {
	return &CustomerOrchestrator{engine}
}

func (or *CustomerOrchestrator) GetByID(ctx context.Context, id uint) (*models.Customer, error) {
	customer, err := or.Engine.DataStore.CustomerRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
func (or *CustomerOrchestrator) Create(ctx context.Context, body *web.Customer) (*models.Customer, error) {
	model := &models.Customer{
		Name:    body.Name,
		Comment: body.Comment,
	}
	var customer *models.Customer
	err := or.Engine.DataStore.WithTransaction(func(tx *gorm.DB) error {
		var err error
		customer, err = or.Engine.DataStore.CustomerRepository.WithTx(tx).Create(ctx, model)
		if err != nil {
			return err
		}

		///////elastic
		body.Id = customer.Model.ID
		err = or.Engine.ElasticData.CustomerRepository.LoadData(ctx, body.ToDTO())
		if err != nil {
			return err
		}
		//////////
		return nil
	})
	if err != nil {
		return nil, err
	}

	return customer, nil
}
func (or *CustomerOrchestrator) DeleteMark(ctx context.Context, id uint) error {
	err := or.Engine.DataStore.CustomerRepository.DeleteMark(ctx, id)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
func (or *CustomerOrchestrator) Update(ctx context.Context, body *web.Customer) (*models.Customer, error) {

	model, err := or.GetByID(ctx, body.Id)
	if err != nil {
		return nil, err
	}
	model.Name = body.Name
	model.Comment = body.Comment

	var customer *models.Customer
	err = or.Engine.DataStore.WithTransaction(func(tx *gorm.DB) error {
		var err error
		customer, err = or.Engine.DataStore.CustomerRepository.WithTx(tx).Update(ctx, model)
		if err != nil {
			return err
		}

		///////elastic
		body.Id = customer.Model.ID
		err = or.Engine.ElasticData.CustomerRepository.LoadData(ctx, body.ToDTO())
		if err != nil {
			return err
		}
		//////////
		return nil
	})

	return customer, nil
}
func (or *CustomerOrchestrator) UpdateElastic(ctx context.Context) error {

	maxID, err := or.Engine.DataStore.CustomerRepository.SelectMaxID(ctx)
	if err != nil {
		return err
	}

	for i := 1; i <= int(maxID); i++ {
		customer, err := or.Engine.DataStore.CustomerRepository.GetByID(ctx, uint(i))
		if err != nil {
			log.Println(err)
			continue
		}
		body := dto.Customer{
			Id:          customer.ID,
			Name:        customer.Name,
			Comment:     customer.Comment,
			DeletedMark: !g.IsEmpty(customer.DeletedAt),
		}
		err = or.Engine.ElasticData.CustomerRepository.LoadData(ctx, &body)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	return nil
}
func (or *CustomerOrchestrator) Search(ctx context.Context, value string) (*model.SearchResponse, error) {
	response, err := or.Engine.ElasticData.CustomerRepository.Search(ctx, value)
	if err != nil {
		return nil, err
	}
	return response, nil
}
