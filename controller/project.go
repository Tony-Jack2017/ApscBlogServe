package controller

import (
	common "ApscBlog/common/model"
	"ApscBlog/model"
	"ApscBlog/model/api"
)

func CreateProjectSVC(req *api.CreateProjectReq) (*common.Response, error) {
	comment := model.Project{}
	err := model.AddProject(&comment)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Create comment type successfully .",
	}, nil
}
func UpdateProjectSVC(req *api.UpdateProjectReq) (*common.Response, error) {
	comment := model.Project{
		ProjectID: req.ProjectID,
	}
	err := model.UpdateProject(&comment)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Update comment type successfully .",
	}, nil
}
func GetProjectListSVC(req *api.GetProjectListReq) (*common.ResponseWithList, error) {
	comment := model.Project{}
	pagination := common.Pagination{
		Current: req.Current,
		Size:    req.Size,
	}
	list, err := model.GetProjectList(&comment, &pagination)
	if err != nil {
		return nil, err
	}
	return &common.ResponseWithList{
		ResponseWithData: common.ResponseWithData{
			Response: common.Response{
				Code:    0,
				Success: true,
				Message: "Get comment type successfully .",
			},
			Data: map[string]interface{}{
				"list": list,
			},
		},
	}, nil
}
