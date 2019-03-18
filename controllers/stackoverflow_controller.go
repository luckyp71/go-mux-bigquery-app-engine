package controllers

import (
	"encoding/json"

	"google.golang.org/appengine"

	s "go-bigquery/services"
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
	json.NewEncoder(res).Encode(questions)
}
