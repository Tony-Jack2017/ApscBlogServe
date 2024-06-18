package model

import (
	"ApscBlog/common/model"
	"ApscBlog/tools"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type ArticleTag struct {
	TagID      int64  `json:"tag_id" bson:"tag_id"`
	TagName    string `json:"tag_name" bson:"tag_name"`
	TagIcon    string `json:"tag_icon" bson:"tag_icon"`
	TagCover   string `json:"tag_cover" bson:"tag_cover"`
	ArticleNum int64  `json:"article_num" bson:"article_num"`
	model.BaseTime
}

func AddArticleTag(tag *ArticleTag) error {
	_, err := articleTagConn.InsertOne(context.TODO(), tag)
	if err != nil {
		return err
	}
	return nil
}
func UpdateArticleTag(tag *ArticleTag) error {
	filter := bson.M{"tag_id": tag.TagID}
	_, err := articleTagConn.UpdateOne(context.TODO(), filter, tag)
	if err != nil {
		return err
	}
	return nil
}
func GetArticleTagList(tag *ArticleTag, pagination *model.Pagination) (*[]ArticleTag, error) {
	var res *[]ArticleTag
	data, err := bson.Marshal(tag)
	if err != nil {
		return nil, err
	}
	list, err := userConn.Find(context.TODO(), data, tools.PagFind(pagination))
	if err != nil {
		return nil, err
	}
	err = list.Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
