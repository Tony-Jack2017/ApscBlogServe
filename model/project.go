package model

import (
	"ApscBlog/common/model"
	"ApscBlog/tools"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type Project struct {
	ProjectID    int64  `json:"project_id"`
	ProjectName  string `json:"project_name"`
	ProjectIcon  string `json:"project_icon"`
	ProjectCover string `json:"project_cover"`
	ProjectDesc  string `json:"project_desc"`
	model.BaseTime
}

func AddProject(project *Project) error {
	_, err := projectConn.InsertOne(context.TODO(), project)
	return err
}
func UpdateProject(project *Project) error {
	filter := bson.M{"project_id": project.ProjectID}
	_, err := userConn.UpdateOne(context.TODO(), filter, project)
	return err
}
func GetProjectList(project *Project, pagination *model.Pagination) (*[]Project, error) {
	var res *[]Project
	data, err := bson.Marshal(project)
	if err != nil {
		return nil, err
	}
	list, errRes := userConn.Find(context.TODO(), data, tools.PagFind(pagination))
	if errRes != nil {
		return nil, err
	}
	err = list.Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
