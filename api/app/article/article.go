package article

import (
	common "ApscBlog/common/model"
	article2 "ApscBlog/controller/article"
	"ApscBlog/model/api/article"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetArticleList(ctx *gin.Context) {
	var req article.GetArticleListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(200, common.Response{
			Code:    100,
			Success: false,
			Message: "Params is error",
		})
		return
	}
	resp, errResp := article2.GetArticleListSVC(&req)
	if errResp != nil {
		ctx.JSON(200, common.Response{
			Code:    100,
			Success: false,
			Message: errResp.Error(),
		})
		return
	}
	ctx.JSON(200, resp)

}
func GetArticle(ctx *gin.Context) {
	articleInfoId, err := strconv.ParseInt(ctx.Param("article_id"), 10, 64)
	if err != nil {
		ctx.JSON(200, common.Response{
			Code:    100,
			Success: false,
			Message: "Params is error",
		})
		return
	}
	resp, errResp := article2.GetArticleSVC(articleInfoId)
	if errResp != nil {
		ctx.JSON(200, common.Response{
			Code:    100,
			Success: false,
			Message: errResp.Error(),
		})
		return
	}
	ctx.JSON(200, resp)

}
func CreateArticle(ctx *gin.Context) {
	var req article.CreateArticleReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(200, common.Response{
			Code:    100,
			Success: false,
			Message: "Params is error",
		})
		return
	}
	resp, errResp := article2.CreateArticleSVC(&req)
	if errResp != nil {
		ctx.JSON(200, common.Response{
			Code:    100,
			Success: false,
			Message: errResp.Error(),
		})
		return
	}
	ctx.JSON(200, resp)
}
