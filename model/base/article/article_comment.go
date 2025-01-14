package article

import (
	"ApscBlog/common/model"
	model2 "ApscBlog/model"
	"ApscBlog/tools"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Comment struct {
	CommentID      int64  `json:"comment_id" bson:"comment_id"`
	SenderName     string `json:"sender_name" bson:"sender_name"`
	SenderEmail    string `json:"sender_email" bson:"sender_email"`
	Content        string `json:"content" bson:"content"`
	Status         string `json:"status" bson:"status"`
	ArticleID      int64  `bson:"article_id" bson:"article_id"`
	model.BaseTime `bson:",inline"`
}

func (c *Comment) TableName() string {
	return "article_comments"
}

func AddComment(comment *Comment) error {
	comment.CreatedAt = model.LocalTime(time.Now())
	_, err := model2.ArticleCommentConn.InsertOne(context.TODO(), comment)
	return err
}
func UpdateComment(comment *Comment) error {
	filter := bson.M{"comment_id": comment.CommentID}
	_, err := model2.ArticleCommentConn.UpdateOne(context.TODO(), filter, comment)
	return err
}
func GetCommentList(articleComment *Comment, pagination *model.Pagination) (*[]Comment, int64, error) {
	var res []Comment
	data := tools.CleanEmptyFields(*articleComment)
	list, errRes := model2.ArticleCommentConn.Find(context.TODO(), data, tools.PagFind(pagination))
	if errRes != nil {
		return nil, 0, errRes
	}
	errDec := list.All(context.TODO(), &res)
	if errDec != nil {
		return nil, 0, errDec
	}
	count, errCount := model2.ArticleCommentConn.CountDocuments(context.Background(), data)
	if errDec != nil {
		return nil, 0, errCount
	}
	return &res, count, nil
}
