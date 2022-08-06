package orchestrator

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"Test_Gorm_Fiber_Elastic/pkg/common/models/web"
	"context"
)

type ICustomerOrchestrator interface {
	GetByID(ctx context.Context, id uint) (*models.Customer, error)
	Create(ctx context.Context, body *web.Customer) (*models.Customer, error)
}
type CustomerOrchestrator struct {
	Engine *Engine
}

func NewCustomerOrchestrator(engine *Engine) ICustomerOrchestrator {
	return &CustomerOrchestrator{engine}
}

func (or CustomerOrchestrator) GetByID(ctx context.Context, id uint) (*models.Customer, error) {
	customer, err := or.Engine.DataStore.CustomerRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
func (or CustomerOrchestrator) Create(ctx context.Context, body *web.Customer) (*models.Customer, error) {
	model := &models.Customer{
		Name:    body.Name,
		Comment: body.Comment,
	}
	customer, err := or.Engine.DataStore.CustomerRepository.Create(ctx, model)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
