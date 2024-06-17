package router

import (
	api "ApscBlog/api/admin"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(admin *gin.RouterGroup) {
	UserRoutes(admin.Group("/user"))
	ArticleRoutes(admin.Group("/article"))
	admin.POST("/comment/create")
	admin.POST("/comment/update")
	admin.GET("/comment/list")
	admin.POST("/project/create")
	admin.POST("/project/update")
	admin.GET("/project/list")

}

func UserRoutes(user *gin.RouterGroup) {
	user.POST("/login", api.AccountLogin)
	user.POST("/create", api.AccountSignUp)
	user.GET("/check", api.CheckUserInfo)
	user.POST("/password/modify", api.ModifyPassword)
	user.POST("/update", api.UpdateUserInfo)
	user.GET("/list", api.GetUserList)
}

func ArticleRoutes(article *gin.RouterGroup) {
	article.POST("/article/create", api.CreateArticle)
	article.POST("/article/update", api.UpdateArticle)
	article.GET("/article/list", api.GetArticleList)
	article.POST("/article/tag/create")
	article.POST("/article/tag/update")
	article.GET("/article/tag/list")
	article.POST("/article/type/create")
	article.POST("/article/type/update")
	article.GET("/article/type/list")
}
