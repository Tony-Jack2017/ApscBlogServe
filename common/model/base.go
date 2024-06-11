package model

type Pagination struct {
	Current int `json:"current" form:"current"`
	Size    int `json:"size" form:"size"`
}
