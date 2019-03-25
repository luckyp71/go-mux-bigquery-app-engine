package controllers

import (
	"encoding/json"

	"google.golang.org/appengine"

	s "go-mux-bigquery-app-engine/services"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPostQuestionsController(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	// Get the current project ID
	projectID := appengine.AppID(ctx)
	param := mux.Vars(req)
	questions := s.GetPostQuestionsService(projectID, param["topic"], ctx)
	res.Header().Set("Content-Type", "application/json")
	response := s.ResponseSuccess(questions)
	json.NewEncoder(res).Encode(response)
}
