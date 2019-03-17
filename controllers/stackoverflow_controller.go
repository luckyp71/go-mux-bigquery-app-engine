package controllers

import (
	"encoding/json"
	service "go-bigquery/services"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPostQuestionsController(res http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)
	questions := service.GetPostQuestionsService(param["projectName"], param["topic"])
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(questions)
}
