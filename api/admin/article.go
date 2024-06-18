package admin

import (
	"ApscBlog/controller"
	"ApscBlog/model/api"
	"ApscBlog/tools"
	"github.com/gin-gonic/gin"
)

func GetArticleList(ctx *gin.Context) {
	var req api.GetArticleListReq
	if !tools.HandleBindReq(ctx, req) {
		return
	}
	err, resp := controller.GetArticleListSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func CreateArticle(ctx *gin.Context) {
	var req api.CreateArticleReq
	if !tools.HandleBindReq(ctx, &req) {
		return
	}
	resp, err := controller.CreateArticleSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func UpdateArticle(ctx *gin.Context) {
}

func CreateArticleType(ctx *gin.Context) {
	var req api.CreateArticleTypeReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.CreateArticleTypeSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func UpdateArticleType(ctx *gin.Context) {
	var req api.UpdateArticleTypeReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.UpdateArticleTypeSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func GetArticleTypeList(ctx *gin.Context) {
	var req api.GetArticleTypeListReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.GetArticleTypeListSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}

func CreateArticleTag(ctx *gin.Context) {
	var req api.CreateArticleTagReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.CreateArticleTagSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func UpdateArticleTag(ctx *gin.Context) {
	var req api.UpdateArticleTagReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.UpdateArticleTagSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func GetArticleTagList(ctx *gin.Context) {
	var req api.GetArticleTagListReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.GetArticleTagListSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
