package router

import "github.com/gin-gonic/gin"

func RegisterCommonRoutes(common *gin.RouterGroup) {
	common.POST("/login_in")
}
