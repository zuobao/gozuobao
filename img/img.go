package img

import (
	"github.com/nfnt/resize"
	"image"
	"time"

	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"github.com/zuobao/gozuobao/logger"
)

type Image struct {
	ImageId int64 `db:"imageid"`
	//	UniqueId   string
	UploadTime time.Time
	Size       int64
	Width      int
	Height     int
	Format     string
	//	Path       string
	Basename string // userId_

	UserId      int64
	StoragePath string `json:"-"`
	StorageId1  string // 目前使用的远程存储是哪个
	StorageId2  string
	StorageId3  string
}

func (i Image) setStorageId(index int, storageId string) {
	switch index {
	case 1:
		i.StorageId1 = storageId
		break
	case 2:
		i.StorageId1 = storageId
		break
	case 3:
		i.StorageId1 = storageId
		break
	}
}

func (i Image) appendStorageId(storageId string) {
	if len(i.StorageId1) == 0 {
		i.StorageId1 = storageId
	}
	if len(i.StorageId2) == 0 {
		i.StorageId2 = storageId
	}
	if len(i.StorageId3) == 0 {
		i.StorageId3 = storageId
	}
}

var JpegOptions = jpeg.Options{Quality: 75}

var thumbnailFunc = resize.InterpolationFunction(resize.NearestNeighbor)

func output(i image.Image, w io.Writer, format string) error {

	switch format {
	case "jpeg":
		return jpeg.Encode(w, i, &JpegOptions)
		break
	case "gif":
		return gif.Encode(w, i, nil)
		break
	case "png":
		return png.Encode(w, i)
	}

	return errors.New("不可识别的文件类型")
}

func WriteJpegFile(i image.Image, fpath string, quality int, overwrite bool) error {
	flag := os.O_CREATE | os.O_WRONLY
	if overwrite {
		flag |= os.O_TRUNC
	} else {
		flag |= os.O_EXCL
	}

	fDst, err := os.OpenFile(fpath, flag, os.FileMode(0666))
	if err != nil {
		return err
	}
	defer fDst.Close()

	opt := jpeg.Options{Quality: quality}

	return jpeg.Encode(fDst, i, &opt)
}

// 生成缩略图的方法
func Thumbnail(maxWidth, maxHeight uint, src io.Reader, dst io.Writer) error {
	i, format, err := image.Decode(src)

	if err != nil {
		return err
	}

	out := resize.Thumbnail(maxWidth, maxHeight, i, thumbnailFunc)

	return output(out, dst, format)
}

func ThumbnailImage(maxWidth, maxHeight uint, src image.Image) image.Image {
	return resize.Thumbnail(maxWidth, maxHeight, src, thumbnailFunc)
}

func ThumbnailImageToFile(maxWidth, maxHeight uint, src image.Image, dst string, format string, overwrite bool) error {
	out := resize.Thumbnail(maxWidth, maxHeight, src, thumbnailFunc)

	flag := os.O_CREATE | os.O_WRONLY
	if overwrite {
		flag |= os.O_TRUNC
	} else {
		flag |= os.O_EXCL
	}

	fDst, err := os.OpenFile(dst, flag, os.FileMode(0666))
	if err != nil {
		return err
	}
	defer fDst.Close()

	return output(out, fDst, format)

}

func ThumbnailFile(maxWidth, maxHeight uint, src, dst string, overwrite bool) error {

	fSrc, err := os.OpenFile(src, os.O_RDONLY, os.FileMode(0666))
	if err != nil {
		return err
	}
	defer fSrc.Close()

	flag := os.O_CREATE | os.O_WRONLY
	if overwrite {
		flag |= os.O_TRUNC
	} else {
		flag |= os.O_EXCL
	}

	fDst, err := os.OpenFile(dst, flag, os.FileMode(0666))
	if err != nil {
		return err
	}
	defer fDst.Close()

	return Thumbnail(maxWidth, maxHeight, fSrc, fDst)
}

func Resize(width, height uint, src io.Reader, dst io.Writer) error {
	i, format, err := image.Decode(src)
	if err != nil {
		return err
	}

	out := resize.Resize(width, height, i, resize.InterpolationFunction(resize.NearestNeighbor))
	switch format {
	case "jpeg":
		return jpeg.Encode(dst, out, &JpegOptions)
		break
	case "gif":
		return gif.Encode(dst, out, nil)
		break
	case "png":
		return png.Encode(dst, out)
	}

	return errors.New("不支持的文件格式: " + format)
}

func GetImage(fpath string, decodeFormat string, scalePercentage int) (image.Image, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var i image.Image

	switch decodeFormat {
	case "png":
		i, err = png.Decode(f)
	case "jpeg", "jpg":
		i, err = jpeg.Decode(f)
	case "gif":
		i, err = gif.Decode(f)
	default:
		i, _, err = image.Decode(f)
	}

	if err != nil {
		return nil, err
	}

	if scalePercentage < 100 && scalePercentage >= 1 {
		//		img.Thumbnail()
		size := i.Bounds().Size()
		i = ThumbnailImage(uint(size.X*scalePercentage/100),
			uint(size.Y*scalePercentage/100), i)
	}

	return i, err
}

// 减小文件尺寸大小
func ReduceSize(fullpath string, size int64, width, height uint ) error {
	var (
		f *os.File
	)

	fi, err := os.Stat(fullpath)
	if err != nil {
		logger.Errorln(err)
		return err
	}

	if fi.Size() < size {
		return nil
	}

	f , err = os.Open(fullpath)
	if err == nil {
		defer func () {
			f.Close()
		}()

		var (
			cfg image.Config
		)
		cfg, _, err = image.DecodeConfig(f)
		if err != nil {
			logger.Errorln(err)
			return err
		}

		if cfg.Width * cfg.Height > (int)(width * height) {
			dst := fullpath + "_reducesize"
			f.Close()
			if nil == ThumbnailFile(width, height, fullpath, dst, true) {
				os.Rename(dst, fullpath)
			}
		}
	}

	return err
}
