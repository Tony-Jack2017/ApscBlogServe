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
type GetUserInfoReq struct {
	UserID string `json:"user_id" form:"form_id"`
}
type ModifyPassword struct {
	UserID      int64  `json:"user_id" form:"user_id"`
	Password    string `json:"password" form:"password"`
	NewPassword string `json:"new_password" form:"new_password"`
}
type UpdateUserInfoReq struct {
	UserID int64 `json:"user_id"`
}
type GetUserListReq struct {
}
