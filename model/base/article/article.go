package article

import (
	"ApscBlog/common/model"
	model2 "ApscBlog/model"
	"ApscBlog/tools"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type ArticleInfo struct {
	ArticleInfoID  int64   `json:"article_info_id" bson:"article_info_id,omitempty,unique"`
	Cover          string  `json:"cover" bson:"cover"`
	Title          string  `json:"title" bson:"title"`
	Description    string  `json:"description" bson:"description"`
	Status         string  `json:"status" bson:"status"` // "Idling" | "Available" | "Pending"
	Tags           []int64 `json:"tags" bson:"tags"`
	TypeID         int64   `json:"type" bson:"type"`
	model.BaseTime `bson:",inline"`
}
type Article struct {
	ArticleInfoID int64  `json:"article_info_id" bson:"article_info_id,omitempty,unique"`
	Content       string `json:"content" bson:"content"`
}

func AddArticle(info *ArticleInfo, article *Article) error {
	info.CreatedAt = model.LocalTime(time.Now())
	_, err := model2.ArticleInfoConn.InsertOne(context.TODO(), info)
	_, err = model2.ArticleConn.InsertOne(context.TODO(), article)
	if err != nil {
		return err
	}
	return nil
}
func UpdateArticle(article *Article) {
}

func GetArticle(articleInfoID int64) (*ArticleInfo, *Article, error) {
	var info ArticleInfo
	var cnt Article
	err := model2.ArticleInfoConn.FindOne(context.TODO(), bson.D{{"article_info_id", articleInfoID}}).Decode(&info)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println("No document found with the given filter")
		} else {
			log.Fatal(err)
		}
		return nil, nil, err
	}
	err = model2.ArticleConn.FindOne(context.TODO(), bson.D{{"article_info_id", articleInfoID}}).Decode(&cnt)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println("No document found with the given filter")
		} else {
			log.Fatal(err)
		}
		return nil, nil, err
	}
	return &info, &cnt, nil
}

func GetArticleList(article *ArticleInfo, pagination *model.Pagination) (*[]ArticleInfo, int64, error) {
	var res []ArticleInfo
	data := tools.CleanEmptyFields(*article)
	list, errRes := model2.ArticleInfoConn.Find(context.TODO(), data, tools.PagFind(pagination))
	if errRes != nil {
		return nil, 0, errRes
	}
	errDec := list.All(context.TODO(), &res)
	if errDec != nil {
		return nil, 0, errDec
	}
	count, errCount := model2.ArticleInfoConn.CountDocuments(context.Background(), data)
	if errDec != nil {
		return nil, 0, errCount
	}
	return &res, count, nil
}
