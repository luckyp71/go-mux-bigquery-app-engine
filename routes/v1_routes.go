package routes

import (
	c "go-bigquery/controllers"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/v1").Subrouter()
	subRouter.HandleFunc("/questions/{topic}", c.GetPostQuestionsController).Methods("GET")
	return router
}
