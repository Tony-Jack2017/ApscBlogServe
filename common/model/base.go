package model

type Pagination struct {
	Current *int `json:"current" form:"current" binding:"required"`
	Size    int  `json:"size" form:"size"  binding:"required"`
}
