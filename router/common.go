package router

import (
	common "ApscBlog/api/common"
	"github.com/gin-gonic/gin"
)

func RegisterCommonRoutes(commonRoute *gin.RouterGroup) {
	commonRoute.POST("/file/upload", common.UploadFileSingle)
}
