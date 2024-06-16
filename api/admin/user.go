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
}

func GetUserList(ctx *gin.Context) {
}
func GetUserInfo(ctx *gin.Context) {
}
func UpdateUserInfo(ctx *gin.Context) {
}
