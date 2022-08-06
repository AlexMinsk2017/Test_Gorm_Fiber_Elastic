package orchestrator

import "Test_Gorm_Fiber_Elastic/pkg/common/services/infra/datastore"

type Engine struct {
	DataStore    *datastore.DataStore
	Orchestrator *Orchestrator
}
type Orchestrator struct {
	Engine *Engine
}

func NewOrchestrator(engine *Engine) *Orchestrator {
	newOrchestrator := Orchestrator{
		Engine: engine,
	}
	return &newOrchestrator
}
