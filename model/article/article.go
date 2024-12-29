package article

import (
	"ApscBlog/common/model"
	model2 "ApscBlog/model"
	"ApscBlog/tools"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleInfo struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty,unique"`
	Cover       string             `json:"cover" bson:"cover"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"` // "Idling" | "Available" | "Pending"
	Tags        []Tag              `json:"tags" bson:"tags"`
	Type        Type               `json:"type" bson:"type"`
	model.BaseTime1
}
type Article struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty,unique"`
	ArticleInfoID primitive.ObjectID `json:"article_info_id" bson:"article_info_id"`
	Content       string             `json:"content" bson:"content"`
}

func AddArticle(info *ArticleInfo, article *Article) error {
	res, err := model2.ArticleInfoConn.InsertOne(context.TODO(), info)
	id, _ := res.InsertedID.(primitive.ObjectID)
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
