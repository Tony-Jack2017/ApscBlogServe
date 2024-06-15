package server

import (
	"ApscBlog/common/config"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime/multipart"
	"path"
)

var MinioClient *minio.Client
var minioErr error

func SetupMinio() {
	MinioClient, minioErr = minio.New(fmt.Sprintf("%s:%d", config.Conf.Minio.Host, config.Conf.Minio.Port), &minio.Options{
		Creds:  credentials.NewStaticV4(config.Conf.Minio.AccessKeyID, config.Conf.Minio.SecretAccessKey, ""),
		Secure: config.Conf.Minio.UseSSL,
	})
	if minioErr != nil {
		fmt.Println("minio Err", minioErr)
	} else {
		fmt.Println("minio pointer", MinioClient)
	}
}

func FileUploader(objectName string, file *multipart.FileHeader) (error, string) {
	bucketName := "apsc-blog"
	ctx := context.Background()
	data, errRead := file.Open()
	contentType := GetFileType(file)
	if errRead != nil {
		return errRead, "Read file fail!!!"
	}
	info, err := MinioClient.PutObject(ctx, bucketName, objectName, data, file.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err, "upload file failed !!!"
	}
	url := fmt.Sprintf("http://%s:%d/%s/%s", config.Conf.Minio.Host, config.Conf.Minio.Port, bucketName, info.Key)
	return nil, url
}

func GetFileType(file *multipart.FileHeader) string {
	extension := path.Ext(file.Filename)
	var contentType string
	switch extension {
	case ".jpg":
		contentType = "image/jpeg"
		break
	case ".png":
		contentType = "image/png"
		break
	case ".gif":
		contentType = "image/gif"
		break
	case ".pdf":
		contentType = "application/pdf"
		break
	case ".svg":
		contentType = "text/svg+xml"
		break
	case ".wav":
		contentType = "audio/x-wav"
		break
	case ".mp3":
		contentType = "audio/mp3"
		break
	case ".woff2":
		contentType = "font/woff2"
		break
	case ".doc":
		contentType = "application/msword"
		break
	case ".html":
		contentType = "text/html"
		break
	case ".css":
		contentType = "text/css"
		break
	case ".js":
		contentType = "text/javascript"
		break
	case ".json":
		contentType = "application/json"
		break
	case ".txt":
		contentType = "text/plain"
		break
	case ".zip":
		contentType = "application/zip"
		break
	default:
		contentType = "application/octet-stream"
		break
	}
	return contentType
}
