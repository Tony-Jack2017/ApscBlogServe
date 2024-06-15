package controller

import (
	common "ApscBlog/common/model"
	"ApscBlog/model"
	"ApscBlog/model/api"
	"go.mongodb.org/mongo-driver/bson"
)

func GetArticleListSVC(req *api.GetArticleListReq) (error, *common.ResponseWithList) {
	err, total := model.GetArticlesCount(bson.D{})
	if err != nil {
		return err, nil
	}
	err, list := model.GetArticles(&common.Pagination{
		Current: req.Current,
		Size:    req.Size,
	}, bson.D{})
	if err != nil {
		return err, nil
	}
	return nil, &common.ResponseWithList{
		ResponseWithData: common.ResponseWithData{
			Response: common.Response{
				Code:    0,
				Success: true,
				Message: "Search the article list successfully !!!",
			},
			Data: list,
		},
		Curren: req.Current,
		Size:   req.Size,
		Total:  total,
	}
}

func CreateArticleSVC(req *api.CreateArticleReq) error {
	articleInfo := &model.ArticleInfo{
		Title: req.Title,
		Cover: req.Cover,
	}
	article := &model.Article{
		Content: req.Content,
	}
	ok, err := model.AddArticle(articleInfo, article)
	if !ok {
		return err
	}
	return nil
}
