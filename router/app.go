package router

import "github.com/gin-gonic/gin"

func RegisterAppRoutes(app *gin.RouterGroup) {
	app.GET("/article/list")
	app.GET("/article/:article_id")

	app.GET("/author/list")
	app.GET("/author/:author_id")

	app.GET("/project/list")
	app.GET("/project/:project_id")

	app.GET("/comment/list")
	app.POST("/comment/push")
}
