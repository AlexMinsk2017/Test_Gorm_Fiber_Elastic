package controllers

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/models/web"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
)

type EOrderRepository interface {
	LoadData(ctx context.Context, dbm *web.Order) error
}

type ElasticOrderRepository struct {
	el *elasticsearch.Client
}

func NewElasticOrderRepository(el *elasticsearch.Client) EOrderRepository {
	return &ElasticOrderRepository{el: el}
}

func (rep ElasticOrderRepository) LoadData(ctx context.Context, dbm *web.Order) error {

	marshalJson, err := json.Marshal(dbm)
	if err != nil {
		log.Print(err)
		return err
	}

	request := esapi.IndexRequest{
		Index:      "ID",
		DocumentID: fmt.Sprintf("%d", dbm.Id),
		Body:       strings.NewReader(string(marshalJson)),
	}
	_, err = request.Do(ctx, rep.el)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
