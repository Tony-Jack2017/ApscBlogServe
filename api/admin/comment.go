package admin

import (
	article2 "ApscBlog/controller/article"
	"ApscBlog/model/api/article"
	"ApscBlog/tools"
	"github.com/gin-gonic/gin"
)

func CreateComment(ctx *gin.Context) {
	var req article.CreateArticleCommentReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := article2.CreateArticleCommentSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func UpdateComment(ctx *gin.Context) {
	var req article.UpdateArticleCommentReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := article2.UpdateArticleCommentSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func GetCommentList(ctx *gin.Context) {
	var req article.GetArticleCommentListReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := article2.GetArticleCommentListSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
