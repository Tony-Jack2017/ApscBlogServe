package tools

import (
	"ApscBlog/common/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PageFind(collection *mongo.Collection, pag *model.Pagination) {
	option := options.Find()
	option.SetSkip(int64((pag.Current - 1) * pag.Size))
	option.SetLimit(int64(pag.Size))
}
