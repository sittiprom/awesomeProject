package controller

import (
	"awesomeProject/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func GetVaccineData(writer http.ResponseWriter, request *http.Request) {
	country := mux.Vars(request)["country"]
	url := "https://covid-api.mmediagroup.fr/v1/vaccines?country=" + country
	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(body))
	data_obj := model.Vaccine{}
	jsonErr := json.Unmarshal(body, &data_obj)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(data_obj)

}
