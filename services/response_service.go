package services

import (
	m "go-mux-bigquery-app-engine/models"
	"net/http"
)

func ResponseSuccess(data interface{}) m.Response {
	resp := m.Response{ResponseCode: http.StatusOK, ResponseDesc: http.StatusText(http.StatusOK), Data: data, DebugInfo: ""}
	return resp
}

func ResponseError(code int, desc string, debugInfo string) m.Response {
	resp := m.Response{ResponseCode: code, ResponseDesc: desc, Data: nil, DebugInfo: debugInfo}
	return resp
}
