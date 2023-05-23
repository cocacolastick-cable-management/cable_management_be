package database

import (
	"github.com/cable_management/cable_management_be/config"
	"github.com/cable_management/cable_management_be/src/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB
)

func init() {

	dsn := config.ENV.DbDsn

	var err error = nil

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	err = DB.AutoMigrate(&entities.Admin{}, &entities.Supplier{}, &entities.Planner{}, &entities.Contractor{})
	if err != nil {
		log.Fatal(err)
		return
	}
}
