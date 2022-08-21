package controllers

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models/dto"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strconv"
	"strings"
)

type ECustomerRepository interface {
	LoadData(ctx context.Context, dbm *dto.Customer) error
	Search(ctx context.Context, value *string) (*esapi.Response, error)
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
func (rep *ElasticCustomerRepository) Search(ctx context.Context, value *string) (*esapi.Response, error) {

	querytype := "prefix" //"match"
	key := "Name"

	var buffer bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			querytype: map[string]interface{}{
				key: *value,
			},
		},
	}
	json.NewEncoder(&buffer).Encode(query)
	response, err := rep.el.Search(rep.el.Search.WithIndex("idcustomer"), rep.el.Search.WithBody(&buffer))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		craft := hit.(map[string]interface{})["_source"].(map[string]interface{})
		fmt.Println(craft)
	}

	return response, nil
}
