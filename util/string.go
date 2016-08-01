package util

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)


var validEmail = regexp.MustCompile(`^.+@.*\..*$`)
var validMobile = regexp.MustCompile(`^1\d{10}$`)
var nonDigitals = regexp.MustCompile(`^\D*$`)

func IsEmail(s string) bool {
	return validEmail.MatchString(s)
}

func IsMobile( s string) bool {
	return validMobile.MatchString(s)
}


// 去掉非数字字符后，再转换成字符串
func ForceToInt64( s string) int64 {
	result := nonDigitals.ReplaceAll([]byte(s), []byte(""))
	if result == nil || len(result) == 0 {
		return 0
	}
	s2 := string(result)
	value, _ := strconv.ParseInt(s2, 10, 0)
	return value
}


func ParseScreen(screen string ) (long int, short int) {
	parts := strings.Split(screen, "x")
	if len(parts) == 2 {
		var err error
		long, err = strconv.Atoi(parts[0])
		if err == nil {
			short , err = strconv.Atoi(parts[1])
			if err == nil {
				if short > long {
					tmp := long
					long = short
					short = tmp
				}
			}
		}
	}

	return long, short
}



func IsEmpty(s string) bool {
	s = strings.TrimSpace(s)
	return s == ""
}


func IsNotEmpty(s string) bool {
	s = strings.TrimSpace(s)
	return s != ""
}


func IsDate(s string) bool {
	s = strings.TrimSpace(s)
	_, err := time.Parse("2006-1-2", s)
	return err == nil
}


func A() {

}
