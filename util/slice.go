package util

import (
	"bytes"
)

func UniqueIntSlice(values []int) []int {
	m := map[int]bool{}
	newSlice := make([]int, 0, len(values))
	if values != nil && len(values) > 0 {
		for _, v := range values {
			if _, existed := m[v]; !existed {
				newSlice = append(newSlice, v)
				m[v] = true
			}
		}
	} else {
		newSlice = values
	}

	return newSlice
}

func UniqueInt64Slice(values []int64) []int64 {
	m := map[int64]bool{}
	newSlice := make([]int64, 0, len(values))
	if values != nil && len(values) > 0 {
		for _, v := range values {
			if _, existed := m[v]; !existed {
				newSlice = append(newSlice, v)
				m[v] = true
			}
		}
	} else {
		newSlice = values
	}

	return newSlice
}

func RepeatString(times int, s, split string) string {
	buf := bytes.NewBufferString("")
	for times > 0 {
		buf.WriteString(s)
		times -= 1
		if times > 0 {
			buf.WriteString(split)
		}
	}
	return buf.String()
}

func StringS2Interface(slice []string) []interface{} {
	interfaces := make([]interface{}, 0)
	for _, s := range slice {
		interfaces = append(interfaces, s)
	}
	return interfaces
}
