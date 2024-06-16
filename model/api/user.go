package api

type AccountLoginReq struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}
type AccountSignUpReq struct {
	Account  string `json:"account" form:"account" binding:"required"`
	Email    string `json:"email" form:"email" binding:"email;required"`
	Code     string `json:"code" form:"code" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
type CheckUserInfoReq struct {
	UserID string `json:"user_id" form:"form_id"`
}
