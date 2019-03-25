package services

import (
	m "go-mux-bigquery-app-engine/models"
)

func ResponseSuccess(data interface{}) m.Response {
	resp := m.Response{ResponseCode: 200, ResponseDesc: "Ok", Data: data, DebugInfo: ""}
	return resp
}

func ResponseError(code int, desc string, debugInfo string) m.Response {
	resp := m.Response{ResponseCode: code, ResponseDesc: desc, Data: nil, DebugInfo: debugInfo}
	return resp
}
