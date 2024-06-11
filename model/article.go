package model

import (
	"ApscBlog/common/model"
	"ApscBlog/model/other"
	"ApscBlog/tools"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleInfo struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty,unique"`
	Cover       string             `json:"cover" bson:"cover"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Tags        []other.ArticleTag `json:"tags"`
	Type        other.ArticleType  `json:"type" bson:"type"`
	Status      string             `json:"status" bson:"status"` // "Idling" | "Available"
	model.BaseTime1
}

type Article struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty,unique"`
	ArticleInfoID int32              `json:"article_info_id"`
	Content       string             `json:"content"`
}

func UpdateArticle() {
}

func AddArticle(info *ArticleInfo, article *Article) (bool, error) {
	_, err := articleConn.InsertOne(context.TODO(), article)
	_, err = articleInfoConn.InsertOne(context.TODO(), info)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetArticles(pag *model.Pagination, filter interface{}) (error, *[]ArticleInfo) {
	var list []ArticleInfo
	cursor, err := articleConn.Find(context.TODO(), filter, tools.PagFind(pag))
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
	count, err := articleConn.CountDocuments(context.TODO(), filter)
	if err != nil {
		return err, 0
	}
	return nil, count
}
