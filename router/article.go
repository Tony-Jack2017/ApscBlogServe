package router

import (
	"ApscBlog/api/app/article"
	"github.com/gin-gonic/gin"
)

func RegisterArticleRoutes(app *gin.RouterGroup) {
	app.GET("/article/list", article.GetArticleList)
	app.GET("/article/:article_id")
	app.POST("/article/add", article.CreateArticle)

	app.GET("/article/comment/list", article.GetArticleCommentList)
	app.POST("/article/comment/push", article.CreateArticleComment)
	app.PUT("/article/comment/edit", article.UpdateArticleComment)

	app.GET("/article/tag/list", article.GetArticleTagList)
	app.POST("/article/tag/add", article.CreateArticleTagList)
	app.PUT("/article/tag/update", article.UpdateArticleTagList)

	app.GET("/article/type/list", article.GetArticleTypeList)
	app.POST("/article/type/add", article.CreateArticleTypeList)
	app.PUT("/article/type/update", article.UpdateArticleTypeList)

	app.GET("/author/list")
	app.GET("/author/:author_id")

	app.GET("/project/list")
	app.GET("/project/:project_id")

}
