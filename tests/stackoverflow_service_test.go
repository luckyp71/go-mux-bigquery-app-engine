package test

import (
	service "go-bigquery/services"
	"testing"
)

func TestGetPostQuestionsController(t *testing.T) {
	expectedResult := true
	actualResult := false

	dataSize := len(service.GetPostQuestionsService("go-mux-ae", "goroutine"))
	if dataSize > 0 {
		actualResult = true
	}

	if expectedResult != actualResult {
		t.Fatalf("Expected %t but got %t", expectedResult, actualResult)
	}
}
