package util

import "strings"



const ApiPrefix = ""
const ConsoleApiPrefix = "/api/c"
const InternalApiPrefix = "/api/i"

func ApiPath(subpath string) string {
	return ApiPrefix + subpath
}

func ConsoleApiPath(subpath string) string {
	return ConsoleApiPrefix + subpath
}

func InternalApiPath(subpath string) string {
	return InternalApiPrefix + subpath
}

func IsApiPath(path string) bool {
	return strings.HasPrefix(path, ApiPrefix)
}



