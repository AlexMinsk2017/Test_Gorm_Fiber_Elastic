package datastore

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/services/datastore/repo"
	"gorm.io/gorm"
)

type DataStore struct {
	dbHandler            *gorm.DB
	CustomerRepository   repo.ICustomerRepository
	GoodRepository       repo.IGoodRepository
	OrderRepository      repo.IOrderRepository
	TableOrderRepository repo.ITableOrderRepository
}

func NewDataStore(dbHandler *gorm.DB) *DataStore {
	dataStore := DataStore{
		dbHandler:            dbHandler,
		CustomerRepository:   repo.NewCustomer(dbHandler),
		GoodRepository:       repo.NewGood(dbHandler),
		OrderRepository:      repo.NewOrder(dbHandler),
		TableOrderRepository: repo.NewTableOrder(dbHandler),
	}
	return &dataStore
}

func (ds *DataStore) WithTransaction(fn func(tx *gorm.DB) error) error {
	return ds.dbHandler.Transaction(fn)
}
