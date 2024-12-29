package article

import (
	common "ApscBlog/common/model"
	article2 "ApscBlog/model/api/article"
	"ApscBlog/model/article"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateArticleSVC(req *article2.CreateArticleReq) (*common.Response, error) {
	articleInfo := &article.ArticleInfo{
		Title:       req.Title,
		Cover:       req.Cover,
		Description: req.Description,
		Status:      "Available",
	}
	cacheArticle := &article.Article{
		Content: req.Content,
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
func GetArticleListSVC(req *article2.GetArticleListReq) (error, *common.ResponseWithList) {
	err, total := article.GetArticlesCount(bson.D{})
	if err != nil {
		return err, nil
	}
	err, list := article.GetArticles(&common.Pagination{
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
