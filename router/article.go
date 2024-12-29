package router

import "github.com/gin-gonic/gin"

func RegisterArticleRoutes(app *gin.RouterGroup) {
	app.GET("/article/list")
	app.GET("/article/:article_id")
	app.GET("/article/comment/list")
	app.POST("/article/comment/push")

	app.GET("/author/list")
	app.GET("/author/:author_id")

	app.GET("/project/list")
	app.GET("/project/:project_id")

}
