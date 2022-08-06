package orchestrator

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"Test_Gorm_Fiber_Elastic/pkg/common/models/web"
	"context"
	"time"
)

type IOrderOrchestrator interface {
	GetByID(ctx context.Context, id uint) (*models.Order, error)
	Create(ctx context.Context, body *web.Order) (*models.Order, error)
}
type OrderOrchestrator struct {
	Engine *Engine
}

func NewOrderOrchestrator(engine *Engine) IOrderOrchestrator {
	return &OrderOrchestrator{engine}
}

func (or OrderOrchestrator) GetByID(ctx context.Context, id uint) (*models.Order, error) {
	dataSet, err := or.Engine.DataStore.OrderRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dataSet, nil
}
func (or OrderOrchestrator) Create(ctx context.Context, body *web.Order) (*models.Order, error) {
	model := &models.Order{
		Number:        body.Number,
		Date:          time.Now(),
		Comment:       body.Comment,
		CustomerRefer: body.CustomerRefer,
	}
	dataSet, err := or.Engine.DataStore.OrderRepository.Create(ctx, model)
	if err != nil {
		return nil, err
	}
	return dataSet, nil
}
