package router

import (
	"ApscBlog/api/app/article"
	"github.com/gin-gonic/gin"
)

func RegisterArticleRoutes(app *gin.RouterGroup) {
	app.GET("/article/list", article.GetArticleList)
	app.GET("/article/create", article.CreateArticle)
	app.GET("/article/:article_id")

	app.GET("/article/comment/list", article.GetArticleCommentList)
	app.POST("/article/comment/push", article.CreateArticleComment)
	app.POST("/article/comment/edit", article.UpdateArticleComment)
	app.GET("/article/tag/list", article.GetArticleTagList)
	app.GET("/article/tag/add", article.CreateArticleTagList)
	app.GET("/article/tag/update", article.UpdateArticleTagList)

	app.GET("/author/list")
	app.GET("/author/:author_id")

	app.GET("/project/list")
	app.GET("/project/:project_id")

}
