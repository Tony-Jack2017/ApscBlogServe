package article

import (
	common "ApscBlog/common/model"
	article2 "ApscBlog/controller/article"
	"ApscBlog/model/api/article"
	"github.com/gin-gonic/gin"
)

func GetArticleTagList(ctx *gin.Context) {
	var req article.GetArticleTagListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(200, common.Response{
			Code:    100,
			Success: false,
			Message: "Params is error",
		})
		return
	}
	resp, errResp := article2.GetArticleTagListSVC(&req)
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

func CreateArticleTagList(ctx *gin.Context) {
	var req article.CreateArticleTagReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(200, common.Response{
			Code:    100,
			Success: false,
			Message: "Params is error",
		})
		return
	}
	resp, errResp := article2.CreateArticleTagSVC(&req)
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

func UpdateArticleTagList(ctx *gin.Context) {
}

func DeleteArticleTagList(ctx *gin.Context) {
}
