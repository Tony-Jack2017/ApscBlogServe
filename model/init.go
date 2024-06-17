package model

import (
	"ApscBlog/common/db"
	"go.mongodb.org/mongo-driver/mongo"
)

var articleConn,
	articleInfoConn,
	articleTagConn,
	articleTypeConn,
	userConn *mongo.Collection

func InitModel() {
	articleConn = db.MgoConn.Collection("articles")
	articleInfoConn = db.MgoConn.Collection("article_infos")
	articleTagConn = db.MgoConn.Collection("article_tags")
	articleTypeConn = db.MgoConn.Collection("article_types")
	userConn = db.MgoConn.Collection("users")
}
