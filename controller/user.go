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
		Account:  req.Account,
		Email:    req.Email,
		Password: req.Password,
	}
	err := model.AddUser(&user)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "User signup successfully!!!",
	}, err
}
func CheckUserInfoSVC(req *api.GetUserInfoReq) (*common.ResponseWithData, error) {
	user := model.User{}
	err, res := model.SearchUser(&user)
	if err != nil {
		return nil, err
	}
	return &common.ResponseWithData{
		Response: common.Response{
			Code:    0,
			Success: true,
			Message: "Get user info successfully .",
		},
		Data: res,
	}, nil
}
func UpdateUserInfoSVC(req *api.UpdateUserInfoReq) (*common.Response, error) {
	user := model.User{}
	err := model.UpdateUser(&user)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Update user info successfully .",
	}, nil
}
func ModifyPasswordSVC(req *api.ModifyPassword) {
}
func GetUserListSVC(req *api.GetUserListReq) (*common.ResponseWithData, error) {
	user := model.User{}
	err, list := model.GetUserList(&user)
	if err != nil {
		return nil, err
	}
	return &common.ResponseWithData{
		Response: common.Response{
			Code:    0,
			Success: true,
			Message: "Get user list successfully .",
		},
		Data: list,
	}, err
}
