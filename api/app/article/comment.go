package article

import (
	common "ApscBlog/common/model"
	article2 "ApscBlog/controller/article"
	"ApscBlog/model/api/article"
	"github.com/gin-gonic/gin"
)

func GetArticleCommentList(ctx *gin.Context) {
	var req article.GetArticleCommentListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(200, common.Response{
			Code:    100,
			Success: false,
			Message: "Params is error",
		})
		return
	}
	resp, errResp := article2.GetArticleCommentListSVC(&req)
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

func CreateArticleComment(ctx *gin.Context) {
	var req article.CreateArticleCommentReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(200, common.Response{
			Code:    100,
			Success: false,
			Message: "Params is error",
		})
		return
	}
	resp, errResp := article2.CreateArticleCommentSVC(&req)
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

func UpdateArticleComment(ctx *gin.Context) {
	var req article.UpdateArticleCommentReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(200, common.Response{
			Code:    100,
			Success: false,
			Message: "Params is error",
		})
		return
	}
	resp, errResp := article2.UpdateArticleCommentSVC(&req)
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

func DeleteArticleComment(ctx *gin.Context) {
}
