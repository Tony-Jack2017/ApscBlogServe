package model

import (
	"ApscBlog/common/model"
	"ApscBlog/tools"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type ArticleType struct {
	TypeID     int64  `json:"type_id" bson:"type_id"`
	TypeName   string `json:"type_name" bson:"type_name"`
	TypeIcon   string `json:"type_icon" bson:"type_icon"`
	TypeCover  string `json:"type_cover" bson:"type_cover"`
	ArticleNum int64  `json:"article_num" bson:"article_num"`
	model.BaseTime
}

func AddArticleType(articleType *ArticleType) error {
	_, err := articleTypeConn.InsertOne(context.TODO(), articleType)
	return err
}
func UpdateArticleType(articleType *ArticleType) error {
	filter := bson.M{"type_id": articleType.TypeID}
	_, err := userConn.UpdateOne(context.TODO(), filter, articleType)
	return err
}
func GetArticleTypeList(articleType *ArticleType, pagination *model.Pagination) (*[]ArticleType, error) {
	var res *[]ArticleType
	data, err := bson.Marshal(articleType)
	if err != nil {
		return nil, err
	}
	list, errRes := userConn.Find(context.TODO(), data, tools.PagFind(pagination))
	if errRes != nil {
		return nil, err
	}
	err = list.Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
