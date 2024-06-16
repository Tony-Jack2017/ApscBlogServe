package api

type AccountLoginReq struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}
type AccountSignUpReq struct {
	Account  string `json:"account" form:"account" binding:"required"`
	Email    string `json:"email" form:"email" binding:"email;required"`
	Password string `json:"password" form:"password" binding:"required"`
}
type CheckUserInfo struct {
	UserID string `json:"user_id" form:"form_id"`
}
