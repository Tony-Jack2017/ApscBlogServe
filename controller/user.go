package controller

import (
	"ApscBlog/common/constant"
	common "ApscBlog/common/model"
	"ApscBlog/model/api"
	"ApscBlog/model/base"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func VerifyUserSVC(req *api.AccountLoginReq) (*common.Response, error) {
	err, _ := base.SearchUser(bson.M{"account": req.Account, "password": req.Password})
	if errors.Is(err, mongo.ErrNoDocuments) {
		return &common.Response{
			Code:    constant.NoFoundErr,
			Success: false,
			Message: "Account or password error !!!",
		}, nil
	}
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Account login successfully .",
	}, nil
}
func UserCreateSVC(req *api.AccountSignUpReq) (*common.Response, error) {
	user := base.User{
		Account:  req.Account,
		Email:    req.Email,
		Password: req.Password,
		Username: req.Account,
	}
	err, res := base.SearchUser(bson.M{"$or": []bson.M{{"account": req.Account}, {"email": req.Email}}})
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}
	if res != nil {
		return &common.Response{
			Code:    constant.IsExistedErr,
			Success: false,
			Message: "Account or email already existed",
		}, nil
	}
	err = base.AddUser(&user)
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
	user := base.User{}
	err, res := base.SearchUser(&user)
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
	user := base.User{}
	err := base.UpdateUser(&user)
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
	user := base.User{}
	err, list := base.GetUserList(&user)
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
