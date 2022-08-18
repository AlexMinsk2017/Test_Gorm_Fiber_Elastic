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

type EGoodRepository interface {
	LoadData(ctx context.Context, dbm *dto.Good) error
}

type ElasticGoodRepository struct {
	el *elasticsearch.Client
}

func NewElasticGoodRepository(el *elasticsearch.Client) EGoodRepository {
	return &ElasticGoodRepository{el: el}
}

func (rep ElasticGoodRepository) LoadData(ctx context.Context, dbm *dto.Good) error {

	marshalJson, err := json.Marshal(dbm)
	if err != nil {
		log.Print(err)
		return err
	}
	documentID := strconv.FormatUint(uint64(dbm.Id), 10)
	request := esapi.IndexRequest{
		Index:      "idgood",
		DocumentID: documentID,
		Body:       strings.NewReader(string(marshalJson)),
	}
	_, err = request.Do(ctx, rep.el)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
