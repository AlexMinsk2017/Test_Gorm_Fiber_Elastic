package controllers

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models/dto"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strconv"
	"strings"
)

type ECustomerRepository interface {
	LoadData(ctx context.Context, dbm *dto.Customer) error
	Search(ctx context.Context, ind string) error
}
type ElasticCustomerRepository struct {
	el *elasticsearch.Client
}

func NewElasticCustomerRepository(el *elasticsearch.Client) ECustomerRepository {
	return &ElasticCustomerRepository{el: el}
}

func (rep *ElasticCustomerRepository) LoadData(ctx context.Context, dbm *dto.Customer) error {

	marshalJson, err := json.Marshal(dbm)
	if err != nil {
		log.Print(err)
		return err
	}

	//documentID := fmt.Sprintf("%d", dbm.Id)
	documentID := strconv.FormatUint(uint64(dbm.Id), 10)

	request := esapi.IndexRequest{
		Index:      "idcustomer",
		DocumentID: documentID,
		Body:       strings.NewReader(string(marshalJson)),
	}
	resp, err := request.Do(ctx, rep.el)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(resp)
	return nil
}
func (rep *ElasticCustomerRepository) Search(ctx context.Context, ind string) error {

	//query := map[string]interface{}{
	//
	//}

	return nil
}
