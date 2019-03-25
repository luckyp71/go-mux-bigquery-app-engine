package main

import (
	r "go-mux-bigquery-app-engine/routes"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	router := r.Routes()
	http.Handle("/", router)
	appengine.Main()
}
