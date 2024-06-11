package admin

import (
	"ApscBlog/controller"
	"ApscBlog/model/api"
	"ApscBlog/tools"
	"github.com/gin-gonic/gin"
)

func GetArticleListInfo(ctx *gin.Context) {
}

func GetArticleList(ctx *gin.Context) {
	var req api.GetArticleListReq
	if !tools.HandleBindReq(ctx, req) {
		return
	}
}

func CreateArticle(ctx *gin.Context) {
	var req api.CreateArticleReq
	if !tools.HandleBindReq(ctx, &req) {
		return
	}
	tools.HandleResponse(ctx, controller.CreateArticleSVC(&req), nil)
}
