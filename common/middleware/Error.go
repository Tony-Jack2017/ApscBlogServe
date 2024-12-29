package middleware

import (
	common "ApscBlog/common/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.JSON(http.StatusInternalServerError, common.Response{
					Code:    500,
					Success: false,
					Message: "Something bad in server, please contact online service .",
				})
			}
		}()
		context.Next()
	}
}
