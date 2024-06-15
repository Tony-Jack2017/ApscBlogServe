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
	err, resp := controller.GetArticleListSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func CreateArticle(ctx *gin.Context) {
	var req api.CreateArticleReq
	if !tools.HandleBindReq(ctx, &req) {
		return
	}
	tools.HandleResponse(ctx, controller.CreateArticleSVC(&req), nil)
}
func UpdateArticle(ctx *gin.Context) {
}
