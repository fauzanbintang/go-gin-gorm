package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	errDb error
)

// connect to DB and models
func Init() {
	dsn := "host=localhost user=postgres password=postgres dbname=orders_by port=5432 sslmode=disable"
	db, errDb = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errDb != nil {
		fmt.Println("Error DB: ", errDb)
	}

	db.Debug().AutoMigrate(&Order{}, &Item{})

	// return db
}

func GetDB() *gorm.DB {
	return db
}
