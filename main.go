package main

import (
	c "go-bigquery/controllers"

	"google.golang.org/appengine"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/questions/{topic}", c.GetPostQuestionsController).Methods("GET")
	http.Handle("/", router)
	appengine.Main()
}
