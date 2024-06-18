package api

import common "ApscBlog/common/model"

type CreateCommentReq struct {
	SenderName    string `json:"sender_name" form:"sender_name" binding:"required"`
	SenderEmail   string `json:"sender_email" form:"sender_email" binding:"required;email"`
	Content       string `json:"content" form:"content" binding:"required"`
	ArticleInfoID int64  `json:"article_info_id" form:"article_info_id" binding:"required"`
}
type UpdateCommentReq struct {
	CommentID int64 `json:"comment_id" binding:"required"`
}
type GetCommentListReq struct {
	common.Pagination
}
