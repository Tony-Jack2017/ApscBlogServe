package article

import (
	"ApscBlog/common/model"
	model2 "ApscBlog/model"
	"ApscBlog/tools"
	"context"
	"time"
)

type ArticleInfo struct {
	ArticleInfoID  int64   `json:"article_info_id" bson:"_id,omitempty,unique"`
	Cover          string  `json:"cover" bson:"cover"`
	Title          string  `json:"title" bson:"title"`
	Description    string  `json:"description" bson:"description"`
	Status         string  `json:"status" bson:"status"` // "Idling" | "Available" | "Pending"
	Tags           []int64 `json:"tags" bson:"tags"`
	TypeID         []int64 `json:"type" bson:"type"`
	model.BaseTime `bson:",inline"`
}
type Article struct {
	ArticleID     int64  `json:"article_id" bson:"_id,omitempty,unique"`
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
func GetArticles(pag *model.Pagination, filter interface{}) (error, *[]ArticleInfo) {
	var list []ArticleInfo
	cursor, err := model2.ArticleConn.Find(context.TODO(), filter, tools.PagFind(pag))
	if err != nil {
		return err, nil
	}
	err = cursor.All(context.TODO(), &list)
	if err != nil {
		return err, nil
	}
	return nil, &list
}
func GetArticlesCount(filter interface{}) (error, int64) {
	count, err := model2.ArticleConn.CountDocuments(context.TODO(), filter)
	if err != nil {
		return err, 0
	}
	return nil, count
}
