package api

import common "ApscBlog/common/model"

type GetArticleListReq struct {
	common.Pagination
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

type CreateArticleTypeReq struct {
	TypeName  string `json:"type_name" form:"type_name"`
	TypeIcon  string `json:"type_icon" form:"type_icon"`
	TypeCover string `json:"type_cover" form:"type_cover"`
}
type UpdateArticleTypeReq struct {
	TypeID    int64  `json:"type_id" form:"type_id" binding:"required"`
	TypeName  string `json:"type_name" form:"type_name"`
	TypeIcon  string `json:"type_icon" form:"type_icon"`
	TypeCover string `json:"type_cover" form:"type_cover"`
}
type GetArticleTypeListReq struct {
	common.Pagination
}

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
	common.Pagination
}
