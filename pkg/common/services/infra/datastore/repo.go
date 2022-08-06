package datastore

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/services/datastore/repo"
	"gorm.io/gorm"
)

type DataStore struct {
	dbHandler  *gorm.DB
	Customer   repo.ICustomerRepository
	Good       repo.IGoodRepository
	Order      repo.IOrderRepository
	TableOrder repo.ITableOrderRepository
}

func NewDataStore(dbHandler *gorm.DB) *DataStore {
	dataStore := DataStore{
		dbHandler:  dbHandler,
		Customer:   repo.NewCustomer(dbHandler),
		Good:       repo.NewGood(dbHandler),
		Order:      repo.NewOrder(dbHandler),
		TableOrder: repo.NewTableOrder(dbHandler),
	}
	return &dataStore
}
