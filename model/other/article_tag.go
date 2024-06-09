package other

import "ApscBlog/common/model"

type ArticleTag struct {
	TagID      int32  `json:"tag_id"`
	TagName    string `json:"tag_name"`
	TagIcon    string `json:"tag_icon"`
	ArticleNum int    `json:"article_num"`
	model.BaseTime
}
