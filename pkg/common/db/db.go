package db

import (
	"Test_Gorm_Fiber_Elastic/pkg/common/config"
	"Test_Gorm_Fiber_Elastic/pkg/common/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(c *config.Config) *gorm.DB {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Customer{}, &models.Good{}, &models.Order{}, &models.TableOrder{})

	return db
}
