package storage

import (
	"io"

	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/zuobao/gozuobao/logger"
	"github.com/zuobao/gozuobao/util"
)

type LocalStorage struct {
	Domain    string
	LocalRoot string
}

func (me *LocalStorage) UploadFile(f io.Reader, length int64, storagePath string) (string, error) {

	fullpath := filepath.Clean(me.LocalRoot + "/" + storagePath)

	err := SaveLocal(f, fullpath)
	if err != nil {
		logger.Errorln(err)
		return "", err
	}

	return me.GetUrl(storagePath), nil
}

func (me *LocalStorage) UploadPublicFile(f io.Reader, length int64, storagePath string) (string, error) {
	return me.UploadFile(f, length, storagePath)
}

func (me *LocalStorage) UploadPrivateImage(f io.Reader, length int64, storagePath string) (string, error) {
	return me.UploadFile(f, length, storagePath)
}
func (me *LocalStorage) UploadPublicImage(f io.Reader, length int64, storagePath string) (string, error) {
	return me.UploadFile(f, length, storagePath)
}

func (me *LocalStorage) GetThumbnailUrl(originalUrl string, width, height int) string {
	return originalUrl + ("-" + strconv.Itoa(width) + "x" + strconv.Itoa(height))
}

func (me *LocalStorage) GetLocalRoot() string {
	return me.LocalRoot
}

func (me *LocalStorage) GetLocalPath(filepath string) string {
	return path.Clean(me.LocalRoot + "/" + filepath)
}

func (me *LocalStorage) GetUrl(storagePath string) string {
	if len(storagePath) > 0 {
		if storagePath[0] == '/' {
			return "http://" + me.Domain + storagePath
		}
	}
	return storagePath
}

func (me *LocalStorage) RemoveFile(path string) error {
	fullpath := me.GetLocalPath(path)
	//log.Println("LocalPath:", fullpath)
	fi, err := os.Stat(fullpath)
	if err == nil && fi.Mode().IsRegular() {
		os.Remove(fullpath)
	} else if err != nil {
		//log.Println(err)
		if os.IsNotExist(err) {
			err = nil
		}
	}
	return err
}

//
//	时长，以秒为单位
//
func (me *LocalStorage) GetPrivateUrlWithToken(uri string, timestamp, userId int64, userSecretKey string) string {
	return uri + "?" + strconv.FormatInt(timestamp, 10)
}

func SaveLocal(in io.Reader, fpath string) error {
	//absolutePath := filepath.Clean(app.AppSetting.LocalStorageRoot + "/" + fpath)
	dir := filepath.Dir(fpath)
	var err error

	if !util.IsDirExists(dir) {

		err = os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}
	}

	out, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer out.Close()

	io.Copy(out, in)

	return nil
}

func (me *LocalStorage) Invalidate(path string) error {
	return nil
}
