package main

import (
	cont "go-bigquery/controllers"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/{projectName}").Subrouter()
	subRouter.HandleFunc("/questions/{topic}", cont.GetPostQuestionsController).Methods("GET")
	http.Handle("/", router)
	appengine.Main()
}
