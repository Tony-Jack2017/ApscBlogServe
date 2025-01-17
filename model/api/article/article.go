package article

import common "ApscBlog/common/model"

type GetArticleListReq struct {
	common.Pagination
	Params interface{}
}
type CreateArticleReq struct {
	Title       string  `json:"title" binding:"required"`
	Cover       string  `json:"cover" binding:"required"`
	TagList     []int64 `json:"tag_list" binding:"required"`
	Type        int64   `json:"type"  binding:"required"`
	Content     string  `json:"content"  binding:"required"`
	Description string  `json:"description" binding:"required"`
}

type UpdateArticleReq struct {
	ArticleInfoID int64  `json:"article_info_id" form:"article_info_id"`
	Cover         string `json:"cover" form:"cover"`
}
type UpdateArticleContentReq struct {
	ArticleID int64  `json:"article_id"`
	Content   string `json:"content"`
}
