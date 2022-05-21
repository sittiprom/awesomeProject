package main

import (
	"awesomeProject/controller"
	"awesomeProject/service"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var DB *gorm.DB

func main() {
	LoadAppConfig()
	service.Connect(AppConfig.ConnectionString)
	service.Migrate()
	router := mux.NewRouter().StrictSlash(true)
	RegisterApiRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))

}

func RegisterApiRoutes(router *mux.Router) {
	router.HandleFunc("/employee", controller.GetAllEmployee).Methods("GET")
	router.HandleFunc("/employee/{id}", controller.GetEmployeeById).Methods("GET")
	router.HandleFunc("/employee/create", controller.CreateEmployee).Methods("POST")
	router.HandleFunc("/employee/update/{id}", controller.UpdateEmployee).Methods("PUT")

}
