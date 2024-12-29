package tools

import (
	"ApscBlog/common/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleBindReq[T interface{}](ctx *gin.Context, req T) bool {
	err := ctx.ShouldBind(req)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Response{
			Code:    404,
			Success: false,
			Message: "params errors",
		})
		return false
	}
	return true
}

func HandleResponse(ctx *gin.Context, err error, resp interface{}) {
	if err != nil {
		panic(err)
		//ctx.JSON(http.StatusInternalServerError, model.Response{
		//	Code:    500,
		//	Success: false,
		//	Message: err.Error(),

		//})
		return
	}
	if resp != nil {
		ctx.JSON(http.StatusOK, resp)
	} else {
		ctx.JSON(http.StatusOK, model.Response{
			Code:    0,
			Success: true,
			Message: "Operation Successfully !!!",
		})
	}
}
