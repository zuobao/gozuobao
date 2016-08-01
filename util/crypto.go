package util

import (
	"crypto/md5"
	"encoding/hex"
)



func Md5(src string) string {
	md5_bytes := md5.Sum([]byte(src))
	return hex.EncodeToString(md5_bytes[:])
}
