package main

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/config"
	"Test_Gorm_Fiber_Elastic/pkg/common/db"
	"Test_Gorm_Fiber_Elastic/pkg/common/services/infra/datastore"
	"Test_Gorm_Fiber_Elastic/pkg/common/services/orchestrator"
	"log"
)

func main() {
	con, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}
	client := db.Init(&con)
	dataStore := datastore.NewDataStore(client)
	engine := orchestrator.Engine{DataStore: dataStore}
	engine.Orchestrator = orchestrator.NewOrchestrator(&engine)

}
