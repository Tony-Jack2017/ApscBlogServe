package controller

import (
	common "ApscBlog/common/model"
	"ApscBlog/model"
	"ApscBlog/model/api"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateArticleSVC(req *api.CreateArticleReq) (*common.Response, error) {
	articleInfo := &model.ArticleInfo{
		Title: req.Title,
		Cover: req.Cover,
	}
	article := &model.Article{
		Content: req.Content,
	}
	err := model.AddArticle(articleInfo, article)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Create article successfully .",
	}, nil
}
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

func CreateArticleTypeSVC(req *api.CreateArticleTypeReq) (*common.Response, error) {
	articleType := model.ArticleType{
		TypeName:  req.TypeName,
		TypeCover: req.TypeCover,
		TypeIcon:  req.TypeIcon,
	}
	err := model.AddArticleType(&articleType)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Create article type successfully .",
	}, nil
}
func UpdateArticleTypeSVC(req *api.UpdateArticleTypeReq) (*common.Response, error) {
	articleType := model.ArticleType{
		TypeID:    req.TypeID,
		TypeName:  req.TypeName,
		TypeCover: req.TypeCover,
		TypeIcon:  req.TypeIcon,
	}
	err := model.UpdateArticleType(&articleType)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Update article type successfully .",
	}, nil
}
func GetArticleTypeListSVC(req *api.GetArticleTypeListReq) (*common.ResponseWithList, error) {
	articleType := model.ArticleType{}
	pagination := common.Pagination{
		Current: req.Current,
		Size:    req.Size,
	}
	list, err := model.GetArticleTypeList(&articleType, &pagination)
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
	}, nil
}

func CreateArticleTagSVC(req *api.CreateArticleTagReq) (*common.Response, error) {
	articleTag := model.ArticleTag{
		TagName:  req.TagName,
		TagCover: req.TagCover,
		TagIcon:  req.TagIcon,
	}
	err := model.AddArticleTag(&articleTag)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Create article type successfully .",
	}, nil
}
func UpdateArticleTagSVC(req *api.UpdateArticleTagReq) (*common.Response, error) {
	articleTag := model.ArticleTag{
		TagID:    req.TagID,
		TagName:  req.TagName,
		TagCover: req.TagCover,
		TagIcon:  req.TagIcon,
	}
	err := model.UpdateArticleTag(&articleTag)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Update article type successfully .",
	}, nil
}
func GetArticleTagListSVC(req *api.GetArticleTagListReq) (*common.ResponseWithList, error) {
	articleTag := model.ArticleTag{}
	pagination := common.Pagination{
		Current: req.Current,
		Size:    req.Size,
	}
	list, err := model.GetArticleTagList(&articleTag, &pagination)
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
	}, nil
}
