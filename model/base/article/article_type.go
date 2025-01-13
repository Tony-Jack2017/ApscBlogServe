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
	TypeID         int64  `json:"type_id" bson:"_id"`
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
func GetArticleTypeList(articleType *Type, pagination *model.Pagination) (*[]Type, error) {
	var res *[]Type
	data, err := bson.Marshal(articleType)
	if err != nil {
		return nil, err
	}
	list, errRes := model2.ArticleCommentConn.Find(context.TODO(), data, tools.PagFind(pagination))
	if errRes != nil {
		return nil, err
	}
	err = list.Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
