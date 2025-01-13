package tools

import (
	"ApscBlog/common/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PagFind(pag *model.Pagination) *options.FindOptions {
	option := options.Find()
	option.SetSkip(int64((pag.Current - 1) * pag.Size))
	option.SetLimit(int64(pag.Size))
	return option
}

/* SQL Fuction */
func InsertWithDefaults(ctx context.Context, collection *mongo.Collection, data interface{}) error {

	//if _, exists := data["created_at"]; !exists {
	//	doc["created_at"] = time.Now()
	//}
	_, err := collection.InsertOne(ctx, data)
	return err
}
