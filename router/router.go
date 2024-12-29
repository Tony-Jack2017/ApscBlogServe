package router

import (
	"ApscBlog/common/config"
	"ApscBlog/common/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func init() {
	router.Use(middleware.CorsMiddleware()).Use(middleware.ErrorMiddleware())
	v1 := router.Group("/api/v1")
	RegisterArticleRoutes(v1.Group("/blog"))
	RegisterAdminRoutes(v1.Group("/admin"))
	RegisterCommonRoutes(v1.Group("/common"))
}

func Run() {
	dns := fmt.Sprintf("%s:%d", config.Conf.Api.Host, config.Conf.Api.Port)
	err := router.Run(dns)
	if err != nil {
		return
	}
}
