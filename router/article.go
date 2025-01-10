package router

import (
	"ApscBlog/api/app/article"
	"github.com/gin-gonic/gin"
)

func RegisterArticleRoutes(app *gin.RouterGroup) {
	app.GET("/article/list", article.GetArticleList)
	app.GET("/article/:article_id")
	app.POST("/article/create", article.CreateArticle)

	app.GET("/article/comment/list", article.GetArticleCommentList)
	app.POST("/article/comment/push", article.CreateArticleComment)
	app.POST("/article/comment/edit", article.UpdateArticleComment)
	app.GET("/article/tag/list", article.GetArticleTagList)
	app.POST("/article/tag/add", article.CreateArticleTagList)
	app.POST("/article/tag/update", article.UpdateArticleTagList)

	app.GET("/author/list")
	app.GET("/author/:author_id")

	app.GET("/project/list")
	app.GET("/project/:project_id")

}
