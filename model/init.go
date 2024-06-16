package model

import (
	"ApscBlog/common/db"
	"go.mongodb.org/mongo-driver/mongo"
)

var articleConn,
	articleInfoConn,
	userConn *mongo.Collection

func InitModel() {
	articleConn = db.MgoConn.Collection("articles")
	articleInfoConn = db.MgoConn.Collection("article_infos")
	userConn = db.MgoConn.Collection("users")
}
