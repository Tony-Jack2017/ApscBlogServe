package router

import (
	"ApscBlog/common/config"
	"ApscBlog/common/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func init() {
	gin.DisableConsoleColor();

  logfile := "/path/to/your/file.txt"
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else if err != nil {
		fmt.Println("检查文件时发生错误:", err)
	} else {
		fmt.Println("文件存在")
	}
  if err != nil {
	  fmt.Println("Could not create log file")
	}
  gin.DefaultWriter = io.MultiWriter(logfile)

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
