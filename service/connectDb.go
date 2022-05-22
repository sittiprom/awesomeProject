package service

import (
	"awesomeProject/model"
	"fmt"
	"github.com/go-redis/redis"
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

func ConnectRedis() {
	redisConnector := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",

		Password: "",

		DB: 0,
	})
	pong, err := redisConnector.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong + " Connect to Redis Successfully ")

}
