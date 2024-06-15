package common

import (
	"ApscBlog/common/model"
	"ApscBlog/controller"
	"ApscBlog/tools"
	"github.com/gin-gonic/gin"
)

func UploadFileSingle(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	err, msg := controller.UploadFileSingleSVC(file)
	tools.HandleResponse(ctx, err, &model.ResponseWithData{
		Response: model.Response{
			Code:    0,
			Success: true,
			Message: "Upload successfully .",
		},
		Data: map[string]string{
			"url": msg,
		},
	})
}
