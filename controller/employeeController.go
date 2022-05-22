package controller

import (
	"awesomeProject/model"
	"awesomeProject/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// @Summary       Create Employee
// @Description
// @Tags  Employee
// @Accept       json
// @Produce      json
// @Param        employee  	body   model.Employee   	true "Add Employee"
// @Success      200  {object} model.Employee
// @Router       /employee/create [post]
func CreateEmployee(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var employee model.Employee
	json.NewDecoder(request.Body).Decode(&employee)
	service.Connector.Create(&employee)
	json.NewEncoder(writer).Encode(employee)
}

// @Summary       Get Employee Data
// @Description   Get Employee By Id
// @Tags  Employee
// @Accept       json
// @Produce      json
// @Param        id  	path   string   	true "1"
// @Success      200  {object}  model.Employee
// @Router       /employee/{id} [get]
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

// @Summary       Get Employee Data
// @Description   Get All Employees
// @Tags  Employee
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Employee
// @Router       /employee [get]
func GetAllEmployee(writer http.ResponseWriter, request *http.Request) {
	var employees []model.Employee
	service.Connector.Find(&employees)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(employees)

}

// @Summary       Update Employee
// @Description   Update Employee
// @Tags  Employee
// @Accept       json
// @Produce      json
// @Param        id  	path   string   	true "1"
// @Param        employee  	body   model.Employee   	true "Add Employee"
// @Success      200  {object} model.Employee
// @Router       /employee/update/{id} [put]
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
