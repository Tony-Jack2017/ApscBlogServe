package tools

import (
	"ApscBlog/common/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

func PagFind(pag *model.Pagination) *options.FindOptions {
	option := options.Find()
	option.SetSkip(int64(*pag.Current * pag.Size))
	option.SetLimit(int64(pag.Size))
	return option
}

/* SQL Fuction */

func CleanEmptyFields(tag interface{}) bson.M {
	result := bson.M{}
	val := reflect.ValueOf(tag)
	typ := reflect.TypeOf(tag)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		bsonTag := fieldType.Tag.Get("bson")

		// 跳过无效的 BSON 标签字段
		if bsonTag == "" || bsonTag == "-" {
			continue
		}

		// 检查字段是否为零值
		if !field.IsZero() {
			result[bsonTag] = field.Interface()
		}
	}

	return result
}

func InsertWithDefaults(ctx context.Context, collection *mongo.Collection, data interface{}) error {

	//if _, exists := data["created_at"]; !exists {
	//	doc["created_at"] = time.Now()
	//}
	_, err := collection.InsertOne(ctx, data)
	return err
}
