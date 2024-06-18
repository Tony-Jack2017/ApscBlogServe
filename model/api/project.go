package api

import common "ApscBlog/common/model"

type CreateProjectReq struct {
	ProjectName string `json:"project_name" form:"project_name" binding:"required"`
	ProjectIcon string `json:"project_icon" form:"project_icon" binding:"required"`
	ProjectDesc string `json:"project_desc" form:"project_desc" binding:"required"`
}
type UpdateProjectReq struct {
	ProjectID int64 `json:"project_id" binding:"required"`
}
type GetProjectListReq struct {
	common.Pagination
}
