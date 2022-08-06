package main

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/config"
	"Test_Gorm_Fiber_Elastic/pkg/common/db"
	"Test_Gorm_Fiber_Elastic/pkg/common/services/infra/datastore"
	"Test_Gorm_Fiber_Elastic/pkg/common/services/io/elastic"
	"Test_Gorm_Fiber_Elastic/pkg/common/services/io/web"
	"Test_Gorm_Fiber_Elastic/pkg/common/services/orchestrator"
	"log"
)

func main() {
	con, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	//postgres
	client := db.Init(&con)
	dataStore := datastore.NewDataStore(client)
	engine := orchestrator.Engine{DataStore: dataStore}
	engine.Orchestrator = orchestrator.NewOrchestrator(&engine)

	//elastic
	_, err = elastic.ClientElastic()
	if err != nil {
		log.Fatalln(err)
		return
	}

	//webservice
	server := &web.WebServices{Orchestrator: engine.Orchestrator}
	err = server.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
