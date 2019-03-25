package services

import (
	m "go-mux-bigquery-app-engine/models"
	"net/http"
)

func ResponseSuccess(data interface{}) m.Response {
	resp := m.Response{ResponseCode: http.StatusOK, ResponseDesc: http.StatusText(http.StatusOK), Data: data, DebugInfo: ""}
	return resp
}

func ResponseNotFound() m.Response {
	resp := m.Response{ResponseCode: http.StatusNotFound, ResponseDesc: http.StatusText(http.StatusNotFound), Data: nil, DebugInfo: "data not found"}
	return resp
}
