package article

import (
	common "ApscBlog/common/model"
	article2 "ApscBlog/model/api/article"
	"ApscBlog/model/base/article"
	"fmt"
	"strconv"
	"time"
)

func CreateArticleTypeSVC(req *article2.CreateArticleTypeReq) (*common.Response, error) {
	str := fmt.Sprintf("%d%s", time.Now().Unix(), "001518")
	typeID, errTrans := strconv.ParseInt(str, 10, 64)
	if errTrans != nil {
		return nil, errTrans
	}
	articleType := article.Type{
		TypeID:    typeID,
		TypeName:  req.TypeName,
		TypeCover: req.TypeCover,
		TypeIcon:  req.TypeIcon,
	}
	err := article.AddArticleType(&articleType)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Create article type successfully .",
	}, nil
}
func UpdateArticleTypeSVC(req *article2.UpdateArticleTypeReq) (*common.Response, error) {
	articleType := article.Type{
		TypeID:    req.TypeID,
		TypeName:  req.TypeName,
		TypeCover: req.TypeCover,
		TypeIcon:  req.TypeIcon,
	}
	err := article.UpdateArticleType(&articleType)
	if err != nil {
		return nil, err
	}
	return &common.Response{
		Code:    0,
		Success: true,
		Message: "Update article type successfully .",
	}, nil
}
func GetArticleTypeListSVC(req *article2.GetArticleTypeListReq) (*common.ResponseWithList, error) {
	articleType := article.Type{}
	pagination := common.Pagination{
		Current: req.Current,
		Size:    req.Size,
	}
	list, total, err := article.GetArticleTypeList(&articleType, &pagination)
	if err != nil {
		return nil, err
	}
	return &common.ResponseWithList{
		ResponseWithData: common.ResponseWithData{
			Response: common.Response{
				Code:    0,
				Success: true,
				Message: "Get article type successfully .",
			},
			Data: map[string]interface{}{
				"list": list,
			},
		},
		Current: *req.Current,
		Size:    req.Size,
		Total:   total,
	}, nil
}
