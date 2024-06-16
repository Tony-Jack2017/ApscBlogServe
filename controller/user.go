package controller

import (
	common "ApscBlog/common/model"
	"ApscBlog/model"
	"ApscBlog/model/api"
)

func VerifyUserSVC(req *api.AccountLoginReq) (*common.ResponseWithData, error) {
	user := model.User{
		Account:  req.Account,
		Password: req.Password,
	}
	err, res := model.SearchUser(&user)
	if err != nil {
		return nil, err
	}
	return &common.ResponseWithData{
		Data: res,
		Response: common.Response{
			Code:    0,
			Success: true,
			Message: "Account login successfully .",
		},
	}, nil
}
func UserCreateSVC(req *api.AccountSignUpReq) (*common.Response, error) {
	user := model.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}
	ok, err := model.AddUser(&user)
	if !ok {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Create user successfully!!!",
	}, err
}
func GetUserInfoSVC() {
}
func CheckUserInfoSVC() {
}
