package storage

import (
	"errors"
	"io"

	"strconv"

	"github.com/zuobao/gozuobao/logger"
	"github.com/zuobao/gozuobao/upyun"
)

type UpYunStorage struct {
	Bucket, Username, Password, Domain string
}

//
//	上传文件，依据一个存储配置对象
//	并返回上传后的完整URL路径
//
func (me *UpYunStorage) UploadFile(f io.Reader, length int64, storagePath string) (string, error) {

	logger.Infoln("bucket", me.Bucket, "username", me.Username, "password", me.Password)

	up := upyun.NewUpYun(me.Bucket, me.Username, me.Password)

	err := up.WriteFile(storagePath, length, f, true)
	if err != nil {
		return "", err
	}

	return me.GetUrl(storagePath), nil
}

func (me *UpYunStorage) UploadPublicFile(f io.Reader, length int64, storagePath string) (string, error) {
	return me.UploadFile(f, length, storagePath)
}

func (me *UpYunStorage) UploadPrivateImage(f io.Reader, length int64, storagePath string) (string, error) {
	return me.UploadFile(f, length, storagePath)
}
func (me *UpYunStorage) UploadPublicImage(f io.Reader, length int64, storagePath string) (string, error) {
	return me.UploadFile(f, length, storagePath)
}

func (me *UpYunStorage) GetThumbnailUrl(originalUrl string, width, height int) string {
	return originalUrl
}

func (me *UpYunStorage) GetLocalRoot() string {
	return ""
}

func (me *UpYunStorage) GetLocalPath(filepath string) string {
	return ""
}

func (me *UpYunStorage) GetUrl(storagePath string) string {
	var root string = ""
	if storagePath[0] != '/' {
		root = "/"
	}

	return "http://" + me.Domain + root + storagePath
}

//
//	时长，以秒为单位
//
func (me *UpYunStorage) GetPrivateUrlWithToken(uri string, timestamp int64, userId int64, userSecretKey string) string {
	return uri + "?" + strconv.FormatInt(timestamp, 10)
}

func (me *UpYunStorage) Invalidate(path string) error {

	return nil
}

func (me *UpYunStorage) RemoveFile(path string) error {
	return errors.New("尚未支持")
}
