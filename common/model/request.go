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
	ResponseWithData
	Total   int64 `json:"total"`
	Current int   `json:"current"`
	Size    int   `json:"size"`
}
