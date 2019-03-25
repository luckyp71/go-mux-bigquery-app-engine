package models

type Response struct {
	ResponseCode int         `json:"responseCode"`
	ResponseDesc string      `json:"responseDesc"`
	Data         interface{} `json:"data"`
	DebugInfo    string      `json:"debugInfo"`
}

func ResponseSuccess(data interface{}) Response {
	resp := Response{ResponseCode: 200, ResponseDesc: "Ok", Data: data, DebugInfo: ""}
	return resp
}
