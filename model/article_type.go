package model

import (
	"ApscBlog/common/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type ArticleType struct {
	TypeID    int32  `json:"type_id"`
	TypeName  string `json:"type_name"`
	TypeIcon  string `json:"type_icon"`
	TypeCover string `json:"type_cover"`
	model.BaseTime
}

func AddType(articleType *ArticleType) error {
	_, err := articleTypeConn.InsertOne(context.TODO(), articleType)
	return err
}
func UpdateType(articleType *ArticleType) error {
	filter := bson.M{"type_id": articleType.TypeID}
	_, err := userConn.UpdateOne(context.TODO(), filter, articleType)
	return err
}
func GetTypeList(articleType *ArticleType) (error, *[]ArticleType) {
	var res *[]ArticleType
	data, err := bson.Marshal(articleType)
	if err != nil {
		return err, nil
	}
	list, errRes := userConn.Find(context.TODO(), data)
	if errRes != nil {
		return err, nil
	}
	err = list.Decode(&res)
	if err != nil {
		return err, nil
	}
	return nil, res
}
