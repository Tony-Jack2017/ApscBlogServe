package other

import "ApscBlog/common/model"

type ArticleType struct {
	TypeID     int32  `json:"type_id"`
	TypeName   string `json:"type_name"`
	TypeIcon   string `json:"type_icon"`
	TypeCover  string `json:"type_cover"`
	ArticleNum int    `json:"article_num"`
	model.BaseTime
}
