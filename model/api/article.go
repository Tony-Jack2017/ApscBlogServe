package api

import "ApscBlog/common/model"

type GetArticleListReq struct {
	model.Pagination
	Params interface{}
}
type CreateArticleReq struct {
	Title   string  `json:"title"`
	Cover   string  `json:"cover"`
	TagList []int32 `json:"tag_list"`
	Type    int32   `json:"type"`
	Content string  `json:"content"`
}
type GetArticleContentReq struct {
	ArticleID     int32 `json:"article_id" form:"article_id"`
	ArticleInfoID int32 `json:"article_info_id" form:"article_info_id"`
}
type UpdateArticleReq struct {
	ArticleInfoID int32  `json:"article_info_id" form:"article_info_id"`
	Cover         string `json:"cover" form:"cover"`
}
type UpdateArticleContentReq struct {
	ArticleID int32  `json:"article_id"`
	Content   string `json:"content"`
}
