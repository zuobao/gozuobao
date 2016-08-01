package util

import (
    "strings"
    "strconv"
    "log"
)

func IsPublicIP(ip string) bool {
	switch {
	case strings.HasPrefix(ip, "127."),
		strings.HasPrefix(ip, "10."),
		strings.HasPrefix(ip, "192.168."):
		return false

	case strings.HasPrefix(ip, "172."):
        parts := strings.Split(ip, ".")
        log.Println(parts)
        if len(parts) == 4 {
            part2, _ := strconv.Atoi(parts[1])
            if part2 >= 16 && part2 <= 32 {
                return false
            }
        } else {
            return false
        }
    }
	return true
}
