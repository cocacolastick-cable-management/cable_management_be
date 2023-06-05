package db

import (
	"github.com/cable_management/cable_management_be/config"
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

//func init() {
//
//	dsn := config.ENV.DbDsn
//
//	var err error = nil
//
//	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic(err)
//		return
//	}
//
//	err = DB.AutoMigrate(&entities.User{})
//	if err != nil {
//		panic(err)
//		return
//	}
//}

func Init() {

	var err error

	dsn := config.ENV.DbDsn

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		return
	}

	DB = DB.Debug()

	err = DB.AutoMigrate(&entities.User{}, &entities.Contract{}, entities.WithDrawRequest{}, entities.WithDrawRequestHistory{}, &entities.Notification{})
	if err != nil {
		panic(err)
		return
	}
}
