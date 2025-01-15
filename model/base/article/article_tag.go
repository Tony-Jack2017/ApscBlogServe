package article

import (
	"ApscBlog/common/model"
	model2 "ApscBlog/model"
	"ApscBlog/tools"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Tag struct {
	TagID          int64  `json:"tag_id" bson:"tag_id"`
	TagName        string `json:"tag_name" bson:"tag_name"`
	TagIcon        string `json:"tag_icon" bson:"tag_icon"`
	TagCover       string `json:"tag_cover" bson:"tag_cover"`
	ArticleNum     int64  `json:"article_num" bson:"article_num"`
	model.BaseTime `bson:",inline"`
}

func AddArticleTag(articleTag *Tag) error {
	articleTag.CreatedAt = model.LocalTime(time.Now())
	_, err := model2.ArticleTagConn.InsertOne(context.TODO(), articleTag)
	if err != nil {
		return err
	}
	return nil
}
func UpdateArticleTag(tag *Tag, action bson.M) error {
	var err error
	filter := bson.M{"tag_id": tag.TagID}
	if action != nil {
		_, err = model2.ArticleTagConn.UpdateOne(context.TODO(), filter, action)
	} else {
		_, err = model2.ArticleTagConn.UpdateOne(context.TODO(), filter, tag)
	}
	if err != nil {
		return err
	}
	return nil
}
func GetArticleTagList(articleTag *Tag, pagination *model.Pagination) (*[]Tag, int64, error) {
	var res []Tag
	var list *mongo.Cursor
	var errRes error
	data := tools.CleanEmptyFields(*articleTag)
	if pagination != nil {
		list, errRes = model2.ArticleTagConn.Find(context.TODO(), data, tools.PagFind(pagination))
	} else {
		list, errRes = model2.ArticleTagConn.Find(context.TODO(), data)
	}
	if errRes != nil {
		return nil, 0, errRes
	}
	errDec := list.All(context.TODO(), &res)
	if errDec != nil {
		return nil, 0, errDec
	}
	count, errCount := model2.ArticleTagConn.CountDocuments(context.Background(), data)
	if errDec != nil {
		return nil, 0, errCount
	}
	return &res, count, nil
}
