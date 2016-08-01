package util

import (
	"testing"
	"fmt"
)



func Test_UniqueIntSlice (t * testing.T) {
	v := []int{1,2,3,3,4,1,2,5}
	v2 := UniqueIntSlice(v)
	fmt.Println(v)
	fmt.Println(v2)

	a := map[int]string {}
	a = nil

	for _, _ = range a {

	}

}


func Test_range_nil (t *testing.T) {
}
