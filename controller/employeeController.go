package controller

import (
	"awesomeProject/model"
	"awesomeProject/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateEmployee(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var employee model.Employee
	json.NewDecoder(request.Body).Decode(&employee)
	service.Connector.Create(&employee)
	json.NewEncoder(writer).Encode(employee)
}

func GetEmployeeById(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	var empFromDB model.Employee
	service.Connector.First(&empFromDB, id)
	if empFromDB.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(" Employee Not Found with " + id)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(empFromDB)
}

func GetAllEmployee(writer http.ResponseWriter, request *http.Request) {
	var employees []model.Employee
	service.Connector.Find(&employees)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(employees)

}
func UpdateEmployee(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	var empFromDB model.Employee
	service.Connector.First(&empFromDB, id)
	if empFromDB.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(" Employee Not Found with " + id)
		return
	}
	json.NewDecoder(request.Body).Decode(&empFromDB)
	service.Connector.Save(&empFromDB)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(empFromDB)

}
