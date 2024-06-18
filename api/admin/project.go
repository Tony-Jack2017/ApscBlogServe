package admin

import (
	"ApscBlog/controller"
	"ApscBlog/model/api"
	"ApscBlog/tools"
	"github.com/gin-gonic/gin"
)

func CreateProject(ctx *gin.Context) {
	var req api.CreateProjectReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.CreateProjectSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func UpdateProject(ctx *gin.Context) {
	var req api.UpdateProjectReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.UpdateProjectSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
func GetProjectList(ctx *gin.Context) {
	var req api.GetProjectListReq
	ok := tools.HandleBindReq(ctx, &req)
	if !ok {
		return
	}
	resp, err := controller.GetProjectListSVC(&req)
	tools.HandleResponse(ctx, err, resp)
}
