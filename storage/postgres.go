package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"loan-rest-api/model"
	"log"
)

var LOAN_DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable"
	loanDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	LOAN_DB = loanDB
}

func Migrate() {
	LOAN_DB.AutoMigrate(&model.Loan{})
}
