package admin

import (
	article2 "ApscBlog/controller/article"
	"ApscBlog/model/api/article"
	"ApscBlog/tools"
	"github.com/gin-gonic/gin"
)

func GetArticleList(ctx *gin.Context) {
	var req article.GetArticleListReq
	if !tools.HandleBindReq(ctx, req) {
		return
	}
	err, resp := article2.GetArticleListSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func CreateArticle(ctx *gin.Context) {
	var req article.CreateArticleReq
	if !tools.HandleBindReq(ctx, &req) {
		return
	}
	resp, err := article2.CreateArticleSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func UpdateArticle(ctx *gin.Context) {
}

func CreateArticleType(ctx *gin.Context) {
	var req article.CreateArticleTypeReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := article2.CreateArticleTypeSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func UpdateArticleType(ctx *gin.Context) {
	var req article.UpdateArticleTypeReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := article2.UpdateArticleTypeSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func GetArticleTypeList(ctx *gin.Context) {
	var req article.GetArticleTypeListReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := article2.GetArticleTypeListSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}

func CreateArticleTag(ctx *gin.Context) {
	var req article.CreateArticleTagReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := article2.CreateArticleTagSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func UpdateArticleTag(ctx *gin.Context) {
	var req article.UpdateArticleTagReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := article2.UpdateArticleTagSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func GetArticleTagList(ctx *gin.Context) {
	var req article.GetArticleTagListReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := article2.GetArticleTagListSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
