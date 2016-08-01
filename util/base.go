package util

import (
	"crypto/rand"
	"encoding/hex"
	uuid "github.com/nu7hatch/gouuid"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func ParseInt(s string, defaultValue int64) int64 {
	value, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return defaultValue
	}
	return value
}

func SplitToIntSlice(s string) []int64 {
	ids := strings.Split(s, ",")
	if len(ids) == 0 {
		return []int64{}
	}

	ints := make([]int64, 0, len(ids))
	for _, v := range ids {
		value := ParseInt(v, 0)
		if value > 0 {
			ints = append(ints, value)
		}
	}

	return ints
}

func Getwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		return strconv.Itoa(filepath.Separator)
	}
	return pwd
}

func NewUUID() *uuid.UUID {

	uuid, err := uuid.NewV4()
	if err != nil {
		return nil
	}
	return uuid
}

func IsDirExists(dir string) bool {
	_, err := os.Stat(dir)
	return err != nil && os.IsExist(err)
}

func MergeValues(src, dst url.Values) {
	for k, v := range src {
		dst[k] = v
	}

}

func Int64SliceToString(s []int64) []string {
	retval := make([]string, 0, len(s))
	for _, v := range s {
		retval = append(retval, strconv.FormatInt(v, 10))
	}
	return retval
}

func IntSliceToString(s []int) []string {
	retval := make([]string, 0, len(s))
	for _, v := range s {
		retval = append(retval, strconv.Itoa(v))
	}
	return retval
}

func StringSliceToInt(s []string) []int {
	retval := make([]int, 0, len(s))
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			//logger.Warnln(err)
		}
		retval = append(retval, i)
	}
	return retval
}

func StringSliceToInt64(s []string) []int64 {
	retval := make([]int64, 0, len(s))
	for _, v := range s {
		i, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
		}
		retval = append(retval, i)
	}
	return retval
}

func Today() time.Time {
	now := time.Now()
	y, m, d := now.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Location())
}

func Yesterday() time.Time {
	return Today().Add(time.Hour * -24)
}

var Location *time.Location = time.Local

func GetCurrYearMonth() (int, time.Time) {

	now := time.Now()
	year := now.Year()
	month := now.Month()

	thetime := time.Date(year, month, 1, 0, 0, 0, 0, Location)
	return year*100 + int(month), thetime
}

func GetYearMonth(t time.Time) int {
	year := t.Year()
	month := t.Month()
	return year*100 + int(month)
}

// 201401, 201402 以此类推
func GetPrevYearMonth() (int, time.Time) {

	now := time.Now()
	prevMonth := now.AddDate(0, -1, 0)
	year := prevMonth.Year()
	month := prevMonth.Month()
	//	time.LoadLocation("Asia/ShangHai")
	prevMonth = time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

	return year*100 + int(month), prevMonth
}

func MonthToRange(month int) (begin, end time.Time, err error) {
	month_string := strconv.Itoa(month)

	begin, err = time.Parse("200601", month_string)
	if err != nil {
		return
	}

	end = begin.AddDate(0, 1, 0)

	return
}

var digits = "0123456789abcdefghijklmnopqrstuvwxyz"

func RandNumberString(length int) string {
	b := make([]byte, length)
	rand.Read(b)

	for i, num := range b {
		b[i] = digits[num%10]
	}

	return string(b)
}

// 生成一个随机位数的字符串
func RandString(bits int) string {
	b := make([]byte, bits)
	rand.Read(b)

	for i, num := range b {
		b[i] = digits[num&31]
	}

	return string(b)
}

// 随机字节，并转换成十六进制小写
func RandHex(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return hex.EncodeToString(b)
}

//生成唯一UUID

func UniqueId() string {
	uuid := NewUUID()
	return hex.EncodeToString(uuid[:])
}

func ParseAddress(remoteAddr string) string {
	ip_port_sep_index := strings.LastIndex(remoteAddr, ":")
	if ip_port_sep_index > 0 {
		remoteAddr = remoteAddr[0:ip_port_sep_index]
	}
	return remoteAddr
}

func GetUserAgent(req *http.Request) string {
	return req.Header.Get("User-Agent")
}

func ToTimestamp(t *time.Time) int64 {
	return t.Unix() * 1000
}

func HasFormValue(req *http.Request, key string) bool {
	if req.Form == nil {
		req.ParseMultipartForm(1024 * 1024)
	}
	_, ok := req.Form[key]
	return ok
}

func String2DateTime(s string, loc *time.Location) *time.Time {
	if loc == nil {
		loc = time.Local
	}
	t, err := time.ParseInLocation("2006-1-2 15:4:5", s, loc)
	if err != nil {
		return nil
	} else {
		return &t
	}
}

func NowDate() time.Time {
	t := time.Now()
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func NowDateString() string {
	t := time.Now()
	return t.Format("2006-1-2")
}

func NowDateTimeString() string {
	t := time.Now()
	return t.Format("2006-1-2 15:04:05")
}

func String2Date(s string, loc *time.Location) *time.Time {
	if loc == nil {
		loc = time.Local
	}
	t, err := time.ParseInLocation("2006-1-2", s, loc)
	if err != nil {
		return nil
	} else {
		return &t
	}
}

//
//file, _ := exec.LookPath(os.Args[0])
//path, _ := filepath.Abs(file)
//fmt.Println(path)
//

func GetPartOfInt64Slice(theSlice []int64, fromKey int, size int) []int64 {
	var dest []int64
	if size <= 0 {
		return dest
	}
	if len(theSlice) <= size{
		return theSlice
	}
	if fromKey >= len(theSlice) {
		return dest
	}
	for key, theInt64 := range theSlice {
		if key >= fromKey {
			if len(dest) <= size{
				dest = append(dest, theInt64)
			}
		}
	}
	return dest
}
func GetPartOfStringSlice(theSlice []string, fromKey int, size int) []string {
	var dest []string
	if size <= 0 {
		return dest
	}
	if len(theSlice) <= size{
		return theSlice
	}
	if fromKey >= len(theSlice) {
		return dest
	}
	for key, theString := range theSlice {
		if key >= fromKey {
			if len(dest) <= size{
				dest = append(dest, theString)
			}
		}
	}
	return dest
}