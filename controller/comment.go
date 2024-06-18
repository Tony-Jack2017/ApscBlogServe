package controller

import (
	common "ApscBlog/common/model"
	"ApscBlog/model"
	"ApscBlog/model/api"
)

func CreateCommentSVC(req *api.CreateCommentReq) (*common.Response, error) {
	comment := model.Comment{
		SenderName:    req.SenderName,
		SenderEmail:   req.SenderEmail,
		ArticleInfoID: req.ArticleInfoID,
		Content:       req.Content,
	}
	err := model.AddComment(&comment)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Create comment type successfully .",
	}, nil
}
func UpdateCommentSVC(req *api.UpdateCommentReq) (*common.Response, error) {
	comment := model.Comment{
		CommentID: req.CommentID,
	}
	err := model.UpdateComment(&comment)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Update comment type successfully .",
	}, nil
}
func GetCommentListSVC(req *api.GetCommentListReq) (*common.ResponseWithList, error) {
	comment := model.Comment{}
	pagination := common.Pagination{
		Current: req.Current,
		Size:    req.Size,
	}
	list, err := model.GetCommentList(&comment, &pagination)
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
