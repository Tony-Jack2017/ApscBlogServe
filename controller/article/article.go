package article

import (
	common "ApscBlog/common/model"
	article2 "ApscBlog/model/api/article"
	"ApscBlog/model/base/article"
	"fmt"
	"strconv"
	"time"
)

func CreateArticleSVC(req *article2.CreateArticleReq) (*common.Response, error) {
	str := fmt.Sprintf("%d%s", time.Now().Unix(), "001518")
	articleID, errTrans := strconv.ParseInt(str, 10, 64)
	if errTrans != nil {
		return nil, errTrans
	}
	articleInfo := &article.ArticleInfo{
		ArticleInfoID: articleID,
		Title:         req.Title,
		Cover:         req.Cover,
		Description:   req.Description,
		Status:        "Available",
	}
	cacheArticle := &article.Article{
		ArticleID: articleID,
		Content:   req.Content,
	}
	err := article.AddArticle(articleInfo, cacheArticle)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Create article successfully .",
	}, nil
}
func GetArticleListSVC(req *article2.GetArticleListReq) (*common.ResponseWithList, error) {
	info := article.ArticleInfo{}
	pagination := common.Pagination{
		Current: req.Current,
		Size:    req.Size,
	}
	list, total, err := article.GetArticleList(&info, &pagination)
	if err != nil {
		return nil, err
	}
	return &common.ResponseWithList{
		ResponseWithData: common.ResponseWithData{
			Response: common.Response{
				Code:    0,
				Success: true,
				Message: "Search the article list successfully !!!",
			},
			Data: list,
		},
		Current: *req.Current,
		Size:    req.Size,
		Total:   total,
	}, nil
}
