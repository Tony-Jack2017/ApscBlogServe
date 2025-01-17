package article

import (
	common "ApscBlog/common/model"
	article2 "ApscBlog/model/api/article"
	"ApscBlog/model/base/article"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"time"
)

type getArticleResp struct {
	article.ArticleInfo
	article.Article
}

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
		Tags:          req.TagList,
		TypeID:        req.Type,
		Status:        "Available",
	}
	cacheArticle := &article.Article{
		ArticleInfoID: articleID,
		Content:       req.Content,
	}
	err := article.AddArticle(articleInfo, cacheArticle)
	err = articleEffect(articleInfo)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Create article successfully .",
	}, nil
}
func GetArticleSVC(articleInfoID int64) (*common.ResponseWithData, error) {
	resInfo, resCnt, err := article.GetArticle(articleInfoID)
	if err != nil {
		return nil, err
	}
	resp := getArticleResp{
		*resInfo,
		*resCnt,
	}
	return &common.ResponseWithData{
		Response: common.Response{
			Code:    0,
			Success: true,
			Message: "Search the article list successfully !!!",
		},
		Data: resp,
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

func articleEffect(info *article.ArticleInfo) error {
	var errTag error
	var errType error
	for _, tagID := range info.Tags {
		errTag = article.UpdateArticleTag(&article.Tag{
			TagID: tagID,
		}, bson.M{
			"$inc": bson.M{"article_num": 1},
		})
	}
	errType = article.UpdateArticleType(&article.Type{
		TypeID: info.TypeID,
	}, bson.M{
		"$inc": bson.M{"article_num": 1},
	})
	if errTag != nil {
		return errTag
	}
	if errType != nil {
		return errType
	}
	return nil
}
