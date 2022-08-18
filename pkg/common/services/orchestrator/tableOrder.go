package orchestrator

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"Test_Gorm_Fiber_Elastic/pkg/common/models/web"
	"context"
)

type ITableOrderOrchestrator interface {
	GetByID(ctx context.Context, id uint) (*models.TableOrder, error)
	Create(ctx context.Context, body *web.TableOrder) (*models.TableOrder, error)
}
type TableOrderOrchestrator struct {
	Engine *Engine
}

func NewTableOrderOrchestrator(engine *Engine) ITableOrderOrchestrator {
	return &TableOrderOrchestrator{engine}
}

func (or *TableOrderOrchestrator) GetByID(ctx context.Context, id uint) (*models.TableOrder, error) {
	dataSet, err := or.Engine.DataStore.TableOrderRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dataSet, nil
}
func (or *TableOrderOrchestrator) Create(ctx context.Context, body *web.TableOrder) (*models.TableOrder, error) {
	model := &models.TableOrder{
		OrderRefer: body.OrderRefer,
		GoodRefer:  body.GoodRefer,
		Price:      body.Price,
		Quantity:   body.Quantity,
	}
	dataSet, err := or.Engine.DataStore.TableOrderRepository.Create(ctx, model)
	if err != nil {
		return nil, err
	}
	return dataSet, nil
}
