package article

import common "ApscBlog/common/model"

type CreateArticleTypeReq struct {
	TypeName  string `json:"type_name" form:"type_name" binding:"required"`
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
	TypeName  string `json:"type_name" form:"type_name"`
	TypeIcon  string `json:"type_icon" form:"type_icon"`
	TypeCover string `json:"type_cover" form:"type_cover"`
	common.Pagination
}
