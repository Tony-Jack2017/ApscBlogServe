package model

import (
	"ApscBlog/common/model"
	"ApscBlog/tools"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type Comment struct {
	CommentID     int64  `json:"comment_id" bson:"comment_id"`
	SenderName    string `json:"sender_name" bson:"sender_name"`
	SenderEmail   string `json:"sender_email" bson:"sender_email"`
	Content       string `json:"content" bson:"content"`
	Status        string `json:"status" bson:"status"`
	ArticleInfoID int64  `bson:"article_info_id" bson:"article_info_id"`
	model.BaseTime
}

func AddComment(comment *Comment) error {
	_, err := commentConn.InsertOne(context.TODO(), comment)
	return err
}
func UpdateComment(comment *Comment) error {
	filter := bson.M{"comment_id": comment.CommentID}
	_, err := userConn.UpdateOne(context.TODO(), filter, comment)
	return err
}
func GetCommentList(comment *Comment, pagination *model.Pagination) (*[]Comment, error) {
	var res *[]Comment
	data, err := bson.Marshal(comment)
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
