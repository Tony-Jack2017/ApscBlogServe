package model

import (
	"ApscBlog/common/db"
	"go.mongodb.org/mongo-driver/mongo"
)

var ArticleConn *mongo.Collection
var ArticleTagConn *mongo.Collection
var ArticleTypeConn *mongo.Collection
var ArticleCommentConn *mongo.Collection
var UserConn *mongo.Collection

func InitModel() {
	ArticleConn = db.MgoConn.Collection("articles")
	ArticleTagConn = db.MgoConn.Collection("article_tags")
	ArticleTypeConn = db.MgoConn.Collection("article_types")
	ArticleCommentConn = db.MgoConn.Collection("article_comments")
	UserConn = db.MgoConn.Collection("users")
}
