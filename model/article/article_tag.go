package article

import (
	"ApscBlog/common/model"
	model2 "ApscBlog/model"
	"ApscBlog/tools"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type Tag struct {
	TagID      int64  `json:"tag_id" bson:"tag_id"`
	TagName    string `json:"tag_name" bson:"tag_name"`
	TagIcon    string `json:"tag_icon" bson:"tag_icon"`
	TagCover   string `json:"tag_cover" bson:"tag_cover"`
	ArticleNum int64  `json:"article_num" bson:"article_num"`
	model.BaseTime
}

func AddArticleTag(tag *Tag) error {
	_, err := model2.ArticleTagConn.InsertOne(context.TODO(), tag)
	if err != nil {
		return err
	}
	return nil
}
func UpdateArticleTag(tag *Tag) error {
	filter := bson.M{"tag_id": tag.TagID}
	_, err := model2.ArticleTagConn.UpdateOne(context.TODO(), filter, tag)
	if err != nil {
		return err
	}
	return nil
}
func GetArticleTagList(tag *Tag, pagination *model.Pagination) (*[]Tag, error) {
	var res *[]Tag
	data, err := bson.Marshal(tag)
	if err != nil {
		return nil, err
	}
	list, err := model2.ArticleTagConn.Find(context.TODO(), data, tools.PagFind(pagination))
	if err != nil {
		return nil, err
	}
	err = list.Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
