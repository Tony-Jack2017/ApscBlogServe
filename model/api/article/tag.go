package article

import common "ApscBlog/common/model"

type CreateArticleTagReq struct {
	TagName  string `json:"tag_name" form:"tag_name"`
	TagIcon  string `json:"tag_icon" form:"tag_icon"`
	TagCover string `json:"tag_cover" form:"tag_cover"`
}
type UpdateArticleTagReq struct {
	TagID    int64  `json:"tag_id" form:"tag_id" binding:"required"`
	TagName  string `json:"tag_name" form:"tag_name"`
	TagIcon  string `json:"tag_icon" form:"tag_icon"`
	TagCover string `json:"tag_cover" form:"tag_cover"`
}
type GetArticleTagListReq struct {
	common.Pagination `json:"-"`
}
