package storage

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

//func (s *StorageSetting) StorageSetting {
//
//}

/**
定义一个远程存储的接口，以便整个应用方便访问，
以后我们可以实现自己的接口，或方便切换各类远程存储
*/
type Storage interface {
	UploadPublicFile(f io.Reader, length int64, storagePath string) (string, error)
	UploadPrivateImage(f io.Reader, length int64, storagePath string) (string, error)

	UploadPublicImage(f io.Reader, length int64, storagePath string) (string, error)

	UploadFile(f io.Reader, length int64, storagePath string) (string, error)

	GetThumbnailUrl(originalUrl string, width, height int) string

	GetUrl(storagePath string) string
	GetLocalPath(filepath string) string
	GetLocalRoot() string

	//
	//	token防盗链URL
	//
	GetPrivateUrlWithToken(uri string, timestamp int64, userId int64, userSecretKey string) string

	Invalidate(path string) error
	RemoveFile(path string) error
}

type QiniuStorage struct {
	Bucket string
}

type StorageSetting struct {
	Type string //one of "upyun", "qiniu"

	Upyun UpYunStorage
	Qiniu QiniuStorage
	Local LocalStorage
}

func UploadFile(srcFilepath string, storagePath string, stg Storage) (string, error) {

	if strings.HasPrefix(srcFilepath, "http") {
		// 处理网络请求
		resp, err := http.Get(srcFilepath)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		return stg.UploadFile(resp.Body, resp.ContentLength, storagePath)

	} else {
		// 打开本地文件
		f, err := os.Open(srcFilepath)
		if err != nil {
			return "", err
		}
		fi, err := f.Stat()
		if err != nil {
			return "", err
		}

		if fi.IsDir() {
			return "", errors.New("不是一个文件")
		}

		defer f.Close()

		contentLength := fi.Size()
		return stg.UploadFile(f, contentLength, storagePath)
	}
}
