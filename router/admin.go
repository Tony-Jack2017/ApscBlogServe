package router

import (
	api "ApscBlog/api/admin"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(admin *gin.RouterGroup) {
	admin.POST("/user/sign_in")
	admin.POST("/user/sign_up")

	admin.POST("/article/create", api.CreateArticle)
	admin.POST("/article/update")
	admin.GET("/article/list")

	admin.POST("/article/tag/create")
	admin.POST("/article/tag/update")
	admin.GET("/article/tag/list")

	admin.POST("/article/type/create")
	admin.POST("/article/type/update")
	admin.GET("/article/type/list")

	admin.POST("/comment/create")
	admin.POST("/comment/update")
	admin.GET("/comment/list")

	admin.POST("/project/create")
	admin.POST("/project/update")
	admin.GET("/project/list")

}
