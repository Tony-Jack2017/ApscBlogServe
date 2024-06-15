package controller

import (
	"ApscBlog/common/server"
	"fmt"
	"mime/multipart"
	"time"
)

func UploadFileSingleSVC(file *multipart.FileHeader) (error, string) {
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	err, msg := server.FileUploader(fileName, file)
	if err != nil {
		return err, "upload file to minio server fail"
	}
	return nil, msg
}
