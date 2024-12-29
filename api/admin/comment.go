package admin

import (
	article2 "ApscBlog/controller/article"
	"ApscBlog/model/api/article"
	"ApscBlog/tools"
	"github.com/gin-gonic/gin"
)

func CreateComment(ctx *gin.Context) {
	var req article.CreateCommentReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := article2.CreateCommentSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func UpdateComment(ctx *gin.Context) {
	var req article.UpdateCommentReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := article2.UpdateCommentSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func GetCommentList(ctx *gin.Context) {
	var req article.GetCommentListReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := article2.GetCommentListSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
