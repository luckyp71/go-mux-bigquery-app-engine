package models

type Response struct {
	ResponseCode int         `json:"responseCode"`
	ResponseDesc string      `json:"responseDesc"`
	Data         interface{} `json:"data"`
	DebugInfo    string      `json:"debugInfo"`
}
