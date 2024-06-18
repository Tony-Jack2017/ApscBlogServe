package router

import (
	api "ApscBlog/api/admin"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(admin *gin.RouterGroup) {
	UserRoutes(admin.Group("/user"))
	ArticleRoutes(admin.Group("/article"))

	admin.POST("/comment/create", api.CreateComment)
	admin.POST("/comment/update", api.UpdateComment)
	admin.GET("/comment/list", api.GetCommentList)

	admin.POST("/project/create", api.CreateProject)
	admin.POST("/project/update", api.UpdateProject)
	admin.GET("/project/list", api.GetProjectList)

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
	article.POST("/create", api.CreateArticle)
	article.POST("/update", api.UpdateArticle)
	article.GET("/list", api.GetArticleList)
	article.POST("/tag/create", api.CreateArticleTag)
	article.POST("/tag/update", api.UpdateArticleTag)
	article.GET("/tag/list", api.GetArticleTagList)
	article.POST("/type/create", api.CreateArticleType)
	article.POST("/type/update", api.UpdateArticleType)
	article.GET("/type/list", api.GetArticleTypeList)
}
