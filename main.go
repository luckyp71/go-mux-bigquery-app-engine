package main

import (
	r "go-bigquery/routes"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	router := r.Routes()
	http.Handle("/", router)
	appengine.Main()
}
