package api

type AccountLoginReq struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}
type AccountSignUpReq struct {
	FullName string `json:"fullName" form:"fullName"`
	Email    string `json:"email" form:"email"`
	Code     string `json:"code" form:"code"`
	Password string `json:"password" form:"password"`
}
type CheckUserInfo struct {
	UserID string `json:"user_id" form:"form_id"`
}
