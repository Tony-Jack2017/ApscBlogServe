package model

import "ApscBlog/common/model"

type Comment struct {
	CommentID   int32  `json:"comment_id"`
	SenderName  string `json:"sender_name"`
	SenderEmail string `json:"sender_email"`
	Content     string `json:"content"`
	model.BaseTime
}
