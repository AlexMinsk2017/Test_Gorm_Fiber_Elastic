package datastore

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/services/io/elastic/controllers"
	"github.com/elastic/go-elasticsearch/v8"
)

type DataElastic struct {
	elHandler            *elasticsearch.Client
	CustomerRepository   controllers.ECustomerRepository
	GoodRepository       controllers.EGoodRepository
	OrderRepository      controllers.EOrderRepository
	TableOrderRepository controllers.ETableOrderRepository
}

func NewElasticDataStore(elHandler *elasticsearch.Client) *DataElastic {
	elasticDataStore := DataElastic{
		elHandler:            elHandler,
		CustomerRepository:   controllers.NewElasticCustomerRepository(elHandler),
		GoodRepository:       controllers.NewElasticGoodRepository(elHandler),
		OrderRepository:      controllers.NewElasticOrderRepository(elHandler),
		TableOrderRepository: controllers.NewElasticTableOrderRepository(elHandler),
	}
	return &elasticDataStore
}
