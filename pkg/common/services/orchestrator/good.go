package orchestrator

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"Test_Gorm_Fiber_Elastic/pkg/common/models/web"
	"context"
	"log"
)

type IGoodOrchestrator interface {
	GetByID(ctx context.Context, id uint) (*models.Good, error)
	Create(ctx context.Context, body *web.Good) (*models.Good, error)
	DeleteMark(ctx context.Context, id uint) error
}
type GoodOrchestrator struct {
	Engine *Engine
}

func NewGoodOrchestrator(engine *Engine) IGoodOrchestrator {
	return &GoodOrchestrator{engine}
}

func (or *GoodOrchestrator) GetByID(ctx context.Context, id uint) (*models.Good, error) {
	dataSet, err := or.Engine.DataStore.GoodRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dataSet, nil
}
func (or *GoodOrchestrator) Create(ctx context.Context, body *web.Good) (*models.Good, error) {
	model := &models.Good{
		Name:    body.Name,
		Comment: body.Comment,
		Price:   body.Price,
	}
	dataSet, err := or.Engine.DataStore.GoodRepository.Create(ctx, model)
	if err != nil {
		return nil, err
	}
	return dataSet, nil
}
func (or *GoodOrchestrator) DeleteMark(ctx context.Context, id uint) error {
	err := or.Engine.DataStore.GoodRepository.DeleteMark(ctx, id)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
