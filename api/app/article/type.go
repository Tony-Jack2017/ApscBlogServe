package article

import (
	common "ApscBlog/common/model"
	article2 "ApscBlog/controller/article"
	"ApscBlog/model/api/article"
	"github.com/gin-gonic/gin"
)

func GetArticleTypeList(ctx *gin.Context) {
	var req article.GetArticleTypeListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(200, common.Response{
			Code:    100,
			Success: false,
			Message: "Params is error",
		})
		return
	}
	resp, errResp := article2.GetArticleTypeListSVC(&req)
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

func CreateArticleTypeList(ctx *gin.Context) {
	var req article.CreateArticleTypeReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(200, common.Response{
			Code:    100,
			Success: false,
			Message: "Params is error",
		})
		return
	}
	resp, errResp := article2.CreateArticleTypeSVC(&req)
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

func UpdateArticleTypeList(ctx *gin.Context) {
}

func DeleteArticleTypeList(ctx *gin.Context) {
}
