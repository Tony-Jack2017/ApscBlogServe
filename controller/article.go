package controller

import (
	model2 "ApscBlog/common/model"
	"ApscBlog/model"
	"ApscBlog/model/api"
)

func GetArticleByCondSVC(req *api.GetArticleListReq) (error, *model2.ResponseWithList) {
	return nil, nil
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
