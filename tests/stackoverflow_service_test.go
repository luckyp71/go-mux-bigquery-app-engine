package tests

import (
	service "go-bigquery/services"
	"net/http"
	"testing"

	"google.golang.org/appengine"
)

func TestGetPostQuestionsController(t *testing.T) {
	var req *http.Request
	ctx := appengine.NewContext(req)
	expectedResult := true
	actualResult := false

	dataSize := len(service.GetPostQuestionsService("go-mux-ae", "goroutine", ctx))
	if dataSize > 0 {
		actualResult = true
	}

	if expectedResult != actualResult {
		t.Fatalf("Expected %t but got %t", expectedResult, actualResult)
	}
}
