package article

import (
	common "ApscBlog/common/model"
	article2 "ApscBlog/model/api/article"
	"ApscBlog/model/article"
)

func CreateArticleCommentSVC(req *article2.CreateArticleCommentReq) (*common.Response, error) {
	comment := article.Comment{
		SenderName:  req.SenderName,
		SenderEmail: req.SenderEmail,
		ArticleID:   req.ArticleID,
		Content:     req.CommentContent,
	}
	err := article.AddComment(&comment)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Create comment type successfully .",
	}, nil
}
func UpdateArticleCommentSVC(req *article2.UpdateArticleCommentReq) (*common.Response, error) {
	comment := article.Comment{
		CommentID: req.CommentID,
	}
	err := article.UpdateComment(&comment)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Update comment type successfully .",
	}, nil
}
func GetArticleCommentListSVC(req *article2.GetArticleCommentListReq) (*common.ResponseWithList, error) {
	comment := article.Comment{}
	pagination := common.Pagination{
		Current: req.Current,
		Size:    req.Size,
	}
	list, err := article.GetCommentList(&comment, &pagination)
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
