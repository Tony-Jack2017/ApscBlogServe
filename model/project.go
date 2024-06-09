package model

import "ApscBlog/common/model"

type Project struct {
	ProjectID    int32  `json:"project_id"`
	ProjectName  string `json:"project_name"`
	ProjectIcon  string `json:"project_icon"`
	ProjectCover string `json:"project_cover"`
	ProjectDesc  string `json:"project_desc"`
	GitLink      string `json:"git_link"`
	model.BaseTime
}
