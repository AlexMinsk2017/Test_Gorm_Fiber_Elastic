package orchestrator

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/services/infra/datastore"
)

type Engine struct {
	DataStore    *datastore.DataStore
	Orchestrator *Orchestrator
	ElasticData  *datastore.DataElastic
}
type Orchestrator struct {
	Engine                 *Engine
	CustomerOrchestrator   ICustomerOrchestrator
	GoodOrchestrator       IGoodOrchestrator
	OrderOrchestrator      IOrderOrchestrator
	TableOrderOrchestrator ITableOrderOrchestrator
	UserOrchestrator       IUserOrchestrator
}

func NewOrchestrator(engine *Engine) *Orchestrator {
	newOrchestrator := Orchestrator{
		Engine:                 engine,
		CustomerOrchestrator:   NewCustomerOrchestrator(engine),
		GoodOrchestrator:       NewGoodOrchestrator(engine),
		OrderOrchestrator:      NewOrderOrchestrator(engine),
		TableOrderOrchestrator: NewTableOrderOrchestrator(engine),
		UserOrchestrator:       NewUserOrchestrator(engine),
	}
	return &newOrchestrator
}
