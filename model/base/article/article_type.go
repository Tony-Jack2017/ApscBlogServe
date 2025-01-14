package article

import (
	"ApscBlog/common/model"
	model2 "ApscBlog/model"
	"ApscBlog/tools"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Type struct {
	TypeID         int64  `json:"type_id" bson:"type_id"`
	TypeName       string `json:"type_name" bson:"type_name"`
	TypeIcon       string `json:"type_icon" bson:"type_icon"`
	TypeCover      string `json:"type_cover" bson:"type_cover"`
	ArticleNum     int64  `json:"article_num" bson:"article_num"`
	model.BaseTime `bson:",inline"`
}

func AddArticleType(articleType *Type) error {
	articleType.CreatedAt = model.LocalTime(time.Now())
	_, err := model2.ArticleTypeConn.InsertOne(context.TODO(), articleType)
	return err
}
func UpdateArticleType(articleType *Type) error {
	filter := bson.M{"type_id": articleType.TypeID}
	_, err := model2.ArticleTypeConn.UpdateOne(context.TODO(), filter, articleType)
	return err
}
func GetArticleTypeList(articleType *Type, pagination *model.Pagination) (*[]Type, int64, error) {
	var res []Type
	data := tools.CleanEmptyFields(*articleType)
	list, errRes := model2.ArticleTypeConn.Find(context.TODO(), data, tools.PagFind(pagination))
	if errRes != nil {
		return nil, 0, errRes
	}
	errDec := list.All(context.TODO(), &res)
	if errDec != nil {
		return nil, 0, errDec
	}
	count, errCount := model2.ArticleTypeConn.CountDocuments(context.Background(), data)
	if errDec != nil {
		return nil, 0, errCount
	}
	return &res, count, nil
}
