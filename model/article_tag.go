package model

import (
	"ApscBlog/common/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type ArticleTag struct {
	TagID         int32  `json:"tag_id"`
	TagName       string `json:"tag_name"`
	TagIcon       string `json:"tag_icon"`
	TagCover      string `bson:"tag_cover"`
	ArticleInfoID int64  `json:"article_info_id"`
	model.BaseTime
}

func AddTag(tag *ArticleTag) error {
	_, err := articleTagConn.InsertOne(context.TODO(), tag)
	if err != nil {
		return err
	}
	return nil
}
func UpdateTag(tag *ArticleTag) error {
	filter := bson.M{"tag_id": tag.TagID}
	_, err := articleTagConn.UpdateOne(context.TODO(), filter, tag)
	if err != nil {
		return err
	}
	return nil
}
func GetTagList(tag *ArticleTag) (error, *[]ArticleTag) {
	var res *[]ArticleTag
	data, err := bson.Marshal(tag)
	if err != nil {
		return err, nil
	}
	list, err := userConn.Find(context.TODO(), data)
	if err != nil {
		return err, nil
	}
	err = list.Decode(&res)
	if err != nil {
		return err, nil
	}
	return nil, nil
}
