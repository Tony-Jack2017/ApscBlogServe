package db

import (
	"ApscBlog/common/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var MgoConn *mongo.Database

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dns := fmt.Sprintf("mongodb://%s:%d", config.Conf.Mongo.Host, config.Conf.Mongo.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dns))
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	MgoConn = client.Database("apsc_blog")
	fmt.Println("Connected to MongoDB!")
}
