package controllers

import (
	"encoding/json"

	s "go-mux-bigquery-app-engine/services"
	"net/http"

	"google.golang.org/appengine"

	"github.com/gorilla/mux"
)

func GetPostQuestionsController(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	// Get the current project ID
	projectID := appengine.AppID(ctx)
	param := mux.Vars(req)
	questions := s.GetPostQuestionsService(projectID, param["topic"], ctx)
	var response interface{}
	res.Header().Set("Content-Type", "application/json")
	if questions == nil {
		response = s.ResponseNotFound()
		json.NewEncoder(res).Encode(response)
	} else {
		response = s.ResponseSuccess(questions)
		json.NewEncoder(res).Encode(response)
	}
}
