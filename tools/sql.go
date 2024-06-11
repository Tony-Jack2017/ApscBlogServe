package tools

import (
	"ApscBlog/common/model"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PagFind(pag *model.Pagination) *options.FindOptions {
	option := options.Find()
	option.SetSkip(int64((pag.Current - 1) * pag.Size))
	option.SetLimit(int64(pag.Size))
	return option
}
