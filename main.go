package main

import (
	"awesomeProject/controller"
	"awesomeProject/docs"
	_ "awesomeProject/docs"
	"awesomeProject/service"
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var DB *gorm.DB

func main() {
	LoadAppConfig()
	service.Connect(AppConfig.ConnectionString)
	service.Migrate()
	service.ConnectRedis()
	router := mux.NewRouter().StrictSlash(true)
	docs.SwaggerInfo.Title = "GO Assignment Project"
	docs.SwaggerInfo.Description = " Create REST API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}
	RegisterApiRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))

}

func RegisterApiRoutes(router *mux.Router) {
	router.HandleFunc("/employee", controller.GetAllEmployee).Methods("GET")
	router.HandleFunc("/employee/{id}", controller.GetEmployeeById).Methods("GET")
	router.HandleFunc("/employee/create", controller.CreateEmployee).Methods("POST")
	router.HandleFunc("/employee/update/{id}", controller.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/covid/vaccines/{country}", controller.GetVaccineData).Methods("GET")
	httpSwagger.URL("swagger.json")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

}
