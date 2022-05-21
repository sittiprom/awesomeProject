package service

import (
	"awesomeProject/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Connector *gorm.DB
var err error

func Connect(connectionString string) {
	Connector, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connect to DB Successfully")
}

func Migrate() {
	Connector.AutoMigrate(&model.Employee{})
	log.Println("DataBase Migrate Conpleted")

}
