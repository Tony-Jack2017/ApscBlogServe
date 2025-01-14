package article

import (
	"ApscBlog/common/model"
	model2 "ApscBlog/model"
	"ApscBlog/tools"
	"context"
	"time"
)

type ArticleInfo struct {
	ArticleInfoID  int64   `json:"article_info_id" bson:"article_info_id,omitempty,unique"`
	Cover          string  `json:"cover" bson:"cover"`
	Title          string  `json:"title" bson:"title"`
	Description    string  `json:"description" bson:"description"`
	Status         string  `json:"status" bson:"status"` // "Idling" | "Available" | "Pending"
	Tags           []int64 `json:"tags" bson:"tags"`
	TypeID         []int64 `json:"type" bson:"type"`
	model.BaseTime `bson:",inline"`
}
type Article struct {
	ArticleID     int64  `json:"article_id" bson:"article_id,omitempty,unique"`
	ArticleInfoID int64  `json:"article_info_id" bson:"article_info_id"`
	Content       string `json:"content" bson:"content"`
}

func AddArticle(info *ArticleInfo, article *Article) error {
	info.CreatedAt = model.LocalTime(time.Now())
	res, err := model2.ArticleInfoConn.InsertOne(context.TODO(), info)
	id, _ := res.InsertedID.(int64)
	article.ArticleInfoID = id
	_, err = model2.ArticleConn.InsertOne(context.TODO(), article)
	if err != nil {
		return err
	}
	return nil
}
func UpdateArticle() {
}
func GetArticleList(article *ArticleInfo, pagination *model.Pagination) (*[]ArticleInfo, int64, error) {
	var res []ArticleInfo
	data := tools.CleanEmptyFields(*article)
	list, errRes := model2.ArticleInfoConn.Find(context.TODO(), data, tools.PagFind(pagination))
	if errRes != nil {
		return nil, 0, errRes
	}
	errDec := list.All(context.TODO(), &res)
	if errDec != nil {
		return nil, 0, errDec
	}
	count, errCount := model2.ArticleInfoConn.CountDocuments(context.Background(), data)
	if errDec != nil {
		return nil, 0, errCount
	}
	return &res, count, nil
}
