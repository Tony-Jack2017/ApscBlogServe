package admin

import (
	"ApscBlog/controller"
	"ApscBlog/model/api"
	"ApscBlog/tools"
	"github.com/gin-gonic/gin"
)

func AccountLogin(ctx *gin.Context) {
	var req api.AccountLoginReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.VerifyUserSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func AccountSignUp(ctx *gin.Context) {
	var req api.AccountSignUpReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.UserCreateSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func CheckUserInfo(ctx *gin.Context) {
	var req api.GetUserInfoReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.CheckUserInfoSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func UpdateUserInfo(ctx *gin.Context) {
	var req api.UpdateUserInfoReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.UpdateUserInfoSVC(&req)
	if err != nil {
		return
	}
	tools.HandleResponse(ctx, err, resp)
}
func ModifyPassword(ctx *gin.Context) {
	var req api.ModifyPassword
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
}
func GetUserList(ctx *gin.Context) {
	var req api.GetUserListReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.GetUserListSVC(&req)
	if err != nil {
		return
	}
	tools.HandleResponse(ctx, err, resp)
}
