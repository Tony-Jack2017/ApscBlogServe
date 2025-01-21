package router

import (
	"ApscBlog/common/config"
	"ApscBlog/common/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func init() {
	// gin.DisableConsoleColor();
    // logfile := "/var/server/blog_server/info.log"
	// file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer file.Close()
	// gin.DefaultWriter := io.MultiWriter(os.Stdout, file)
	router.Use(middleware.CorsMiddleware()).Use(middleware.ErrorMiddleware())
	v1 := router.Group("/v1")
	RegisterArticleRoutes(v1.Group("/blog"))
	//RegisterAdminRoutes(v1.Group("/admin"))
	RegisterCommonRoutes(v1.Group("/common"))
}

func Run() {
	dns := fmt.Sprintf("%s:%d", config.Conf.Api.Host, config.Conf.Api.Port)
	err := router.Run(dns)
	if err != nil {
		return
	}
}
