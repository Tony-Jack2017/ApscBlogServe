package article

import common "ApscBlog/common/model"

type CreateArticleCommentReq struct {
	ArticleID      int64  `json:"article_id" form:"article_id" binding:"required"`
	CommentReplyID int64  `json:"comment_reply_id" form:"comment_reply_id"`
	CommentContent string `json:"comment_content" form:"comment_content"`
	SenderName     string `bson:"sender_name" form:"sender_name"`
	SenderEmail    string `bson:"sender_email" form:"sender_email"`
}
type UpdateArticleCommentReq struct {
	CommentID      int64  `json:"comment_id" form:"comment_id" binding:"required"`
	CommentContent string `json:"comment_content" form:"comment_content"`
}
type GetArticleCommentListReq struct {
	ArticleID string `json:"article_id" form:"article_id" binding:"required"`
	common.Pagination
}
