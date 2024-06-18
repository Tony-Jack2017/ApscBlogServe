package admin

import (
	"ApscBlog/controller"
	"ApscBlog/model/api"
	"ApscBlog/tools"
	"github.com/gin-gonic/gin"
)

func CreateComment(ctx *gin.Context) {
	var req api.CreateCommentReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.CreateCommentSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func UpdateComment(ctx *gin.Context) {
	var req api.UpdateCommentReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.UpdateCommentSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func GetCommentList(ctx *gin.Context) {
	var req api.GetCommentListReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.GetCommentListSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
