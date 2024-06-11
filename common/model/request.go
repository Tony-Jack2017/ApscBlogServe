package model

type Response struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type ResponseWithData struct {
	Response
	Data interface{} `json:"data"`
}
type ResponseWithList struct {
	Total  int64 `json:"total"`
	Curren int   `json:"curren"`
	Size   int   `json:"size"`
	ResponseWithData
}
