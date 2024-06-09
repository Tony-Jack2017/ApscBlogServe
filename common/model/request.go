package model

type ReqPagination struct {
	Current int `json:"current" form:"current"`
	Size    int `json:"size" form:"size"`
}

type Response struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ResponseWithData struct {
	Response
	Data interface{} `json:"data"`
}
