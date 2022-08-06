package elastic

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/config"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

func ClientElastic() (*elasticsearch.Client, error) {
	con, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(con)

	cfg := elasticsearch.Config{
		Addresses: []string{con.Addresses},
		Username:  con.Username,
		Password:  con.Password,
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	info, err := client.Info()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer info.Body.Close()
	log.Println(info)

	return client, nil
}
