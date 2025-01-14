package article

import (
	"ApscBlog/common/constant"
	common "ApscBlog/common/model"
	article2 "ApscBlog/model/api/article"
	"ApscBlog/model/base/article"
	"fmt"
	"strconv"
	"time"
)

func CreateArticleTagSVC(req *article2.CreateArticleTagReq) (*common.Response, error) {
	str := fmt.Sprintf("%d%s", time.Now().Unix(), "001518")
	tagID, errTrans := strconv.ParseInt(str, 10, 64)
	if errTrans != nil {
		return nil, errTrans
	}
	articleTag := article.Tag{
		TagID:      tagID,
		TagName:    req.TagName,
		TagCover:   req.TagCover,
		TagIcon:    req.TagIcon,
		ArticleNum: 0,
	}
	err := article.AddArticleTag(&articleTag)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    constant.ActionSuc,
		Success: true,
		Message: "Create article tag successfully .",
	}, nil
}
func UpdateArticleTagSVC(req *article2.UpdateArticleTagReq) (*common.Response, error) {
	articleTag := article.Tag{
		TagID:    req.TagID,
		TagName:  req.TagName,
		TagCover: req.TagCover,
		TagIcon:  req.TagIcon,
	}
	err := article.UpdateArticleTag(&articleTag)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Update article type successfully .",
	}, nil
}
func GetArticleTagListSVC(req *article2.GetArticleTagListReq) (*common.ResponseWithList, error) {
	articleTag := article.Tag{}
	pagination := common.Pagination{
		Current: req.Current,
		Size:    req.Size,
	}
	list, total, err := article.GetArticleTagList(&articleTag, &pagination)
	if err != nil {
		return nil, err
	}
	return &common.ResponseWithList{
		ResponseWithData: common.ResponseWithData{
			Response: common.Response{
				Code:    0,
				Success: true,
				Message: "Get article type successfully .",
			},
			Data: map[string]interface{}{
				"list": list,
			},
		},
		Current: *req.Current,
		Size:    req.Size,
		Total:   total,
	}, nil
}
