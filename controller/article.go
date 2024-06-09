package controller

import (
	"ApscBlog/model"
	"ApscBlog/model/api"
)

func GetArticleByCondSVC() {
}

func CreateArticleSVC(req api.CreateArticleReq) error {

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
