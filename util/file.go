package util

import (
	"io"
	"mime/multipart"
	"os"
)

func IsFileExisted(thepath string) bool {
	stat, err := os.Stat(thepath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return stat.Mode().IsRegular()
}

func IsRegularOrNotExists(thepath string) bool {
	stat, err := os.Stat(thepath)
	if err != nil {
		return os.IsNotExist(err)
	}
	return stat.Mode().IsRegular()
}

func WriteUploadFile(fileheaer *multipart.FileHeader, dst string) error {
	srcFile, err := fileheaer.Open()
	if err != nil {
		return err
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		srcFile.Close()
		return err
	}

	io.Copy(dstFile, srcFile)
	dstFile.Sync()

	dstFile.Close()
	srcFile.Close()

	return nil
}
