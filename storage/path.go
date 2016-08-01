package storage

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
	"regexp"
)

//	为以下资源命名：
//
//		用户头像、女神生活照、女神付费照片
//
//
//
//		系统资源: 礼物图标
//
//	聊天传输，图片，语音
//
//

//var randSource rand.Source
var r *rand.Rand
var r_locker sync.Mutex
var timeBaseline time.Time


var MAGIC_NUMBER int64 = 2147483647

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
	timeBaseline, _ = time.Parse("2006-1-2", "2014-3-21") //2013-3-21 这一天确定用户头像使用不同的图标路径
}

func randomNumber() int64 {
	r_locker.Lock()
	n := r.Int63()
	r_locker.Unlock()
	return n
}

func randomString() string {
	n := randomNumber()
	return strconv.FormatInt(n, 36)

}

func randomNumber32() int32 {
	r_locker.Lock()
	n := r.Int31()
	r_locker.Unlock()
	return n
}

func randomStringShort() string {
	n := randomNumber32()
	return strconv.FormatInt(int64(n), 36)
}

func IsUserIconPath(path string) bool {
	return strings.HasPrefix(path, "/s/icon/u/")
}

func IsVirtualPath(path string) bool {
	return strings.HasPrefix(path, "/s/icon/u/") || strings.HasPrefix(path, "/s/guess/") || strings.HasPrefix(path, "/s/puzzle/")
}

func IsPrivatePath(path string) bool {
	return strings.HasPrefix(path, "/s/pri/")
}

func IsGoddessSpecialImagePath(path string) bool {
	return strings.HasPrefix(path, "/s/gds/")
}


var UserIconSuffixSep = "_"

func UserIconSuffix(storagePath string) string {
	sepindex := strings.LastIndex(storagePath, UserIconSuffixSep)
	if sepindex >= 0 {
		return storagePath[sepindex:]
	} else {
		return ""
	}

}

var magicSuffixRegex = regexp.MustCompile("^[0-9a-zA-Z]+")
func VirtualPathToStoragePath(urlPath string) string {
	if urlPath[0] == '/' {

		sepindex := strings.LastIndex(urlPath, UserIconSuffixSep)

		if sepindex > 0 {
			suffix := urlPath[sepindex+1:]
//			logger.Debugln("suffix:", suffix)
			if len(suffix) > 0 {
				magicSuffix := magicSuffixRegex.FindString(suffix)
//				logger.Debugln("magicSuffix:", magicSuffix)

				suffix = suffix[len(magicSuffix) :]
				urlPath = urlPath[0:sepindex] + suffix
			} else {
				urlPath =  urlPath[0:sepindex]
			}
		}
	}
//	logger.Debugln(urlPath)
	return urlPath
}



func GetGiftIconStoragePath(giftId int64, ) string {
	return "/s/icon/g/" + strconv.FormatInt(giftId, 10) + ".png"
}

func GetItemIconStoragePath(itemId int64) string {
	return "/s/icon/items/" + strconv.FormatInt(itemId, 10) + ".png"
}

func GetBuyLogIcon(buyType int) string {
	return "/s/icon/consume/" + strconv.Itoa(buyType) + ".png"
}



//	生成礼物图标路径，暂时不用
//func GiftIconPath(giftId int64) (localStoragePath string, urlPath string) {
//	// 0x2000 == 8191
////	dir := giftId & 0x1FFF
//	t_part := strconv.FormatInt(int64(time.Now().Sub(timeBaseline)/time.Second), 36)
//	r_part := strconv.FormatInt(r.Int63n(36*36-1), 36)
//	suffix := t_part + r_part
//	localStoragePath = "/s/icon/g/" + strconv.FormatInt(giftId, 10)
//
//	return localStoragePath, localStoragePath + UserIconSuffixSep + suffix
//}


//	用户头像图片
func UserIconPath(userId int64) (localStoragePath string, urlPath string) {
	// 0x2000 == 8191
	dir := userId & 0x1FFF

	t_part := strconv.FormatInt(int64(time.Now().Sub(timeBaseline)/time.Second), 36)
	r_part := strconv.FormatInt(r.Int63n(36*36-1), 36)
	suffix := t_part + r_part

	localStoragePath = "/s/icon/u/" + strconv.FormatInt(dir, 10) + "/" + strconv.FormatInt(userId, 10)

	return localStoragePath, localStoragePath + UserIconSuffixSep + suffix
}

//// 	女神头像图片
//func GoddessIconStoragePath(userId int64) string {
//	// 0x2000 == 8191
//	dir := userId & 0x1FFF
//	return "/s/icon/g/" + strconv.FormatInt(dir, 10) + "/" + strconv.FormatInt(userId, 10)
//}
//
//

func GoddessImagePath(goddessId int64, special string) (localStoragePath string, urlPath string) {
	dir := goddessId & 0x1FFF

	t_part := strconv.FormatInt(int64(time.Now().Sub(timeBaseline)/time.Second), 36)
	r_part := strconv.FormatInt(r.Int63n(36*36-1), 36)
	suffix := t_part + r_part

	localStoragePath = "/s/ns/" + strconv.FormatInt(dir, 10) + "/" + strconv.FormatInt(goddessId, 10) + "_" + special
	return localStoragePath, localStoragePath + UserIconSuffixSep + suffix
}



func CreatePhotoName(userId int64) string {
	src := "user_" + strconv.FormatInt(userId, 36) + "_time_" + strconv.FormatInt(time.Now().UnixNano(), 36)
	bytes := md5.Sum([]byte(src))
	return hex.EncodeToString(bytes[:])
}

// 	/s/v/YYMMDD/goddessId/roomId/timestamp_userid_64bitRANDOM
//	/s/i/YYMMDD/goddessId/roomId/timestamp_userid_64bitRANDOM
//	/s/f/YYMMDD/goddessId/roomId/timestamp_userid_64bitRANDOM
const (
	ChatFileTypeImage = byte('i')	//聊天中的图片
	ChatFileTypeVoice = byte('v')	//聊天的语音
	ChatFileTypeFile  = byte('f')	//其他类文件
)

func ChatFilepath(filetype byte, userId, goddessId, roomId int64, t time.Time) string {

	paths := make([]string, 0, 7)

	paths = append(paths, "/s/c/"+string(filetype), t.Format("20060102"), strconv.FormatInt(goddessId, 10),
		strconv.FormatInt(roomId, 10),
		strconv.FormatInt(t.Unix(), 10)+"_"+strconv.FormatInt(userId, 10)+"_"+randomString())
	thepath := strings.Join(paths, "/")
	return thepath
}


func PhotoFilepath(photoId, userId int64, t time.Time, private bool) (basename string, fpath string)  {
	prefix := "pub"
	if private {
		prefix = "pri"
	}

	names := make([]string, 0, 7)

	randomNumber := randomNumber32()
	var pinId int64 = userId //2147483647
	pinId ^= MAGIC_NUMBER//
	pinId ^= int64(randomNumber)
	pinId ^= photoId

	datePart := t.Format("20060102")

	names = append(names,
		strconv.FormatInt(pinId, 10),
		strconv.FormatInt(photoId, 10),
		strconv.FormatInt(int64(randomNumber), 10))

	basename = strings.Join(names, "_")

	fpath = "/s/"+ prefix +"/" + datePart + "/" + basename

	return basename, fpath

}


func PublicPhotoFilepath(imgId, userId int64, t time.Time) (basename string, fpath string) {
	return PhotoFilepath(imgId, userId, t, false)
}


func PrivatePhotoFilepath(imgId, userId int64, t time.Time) (basename string, fpath string) {
	return PhotoFilepath(imgId, userId, t, true)
}
