package util

import (
	"github.com/martini-contrib/render"
	"math/rand"
)




func RenderHtml(r render.Render, view string, data interface {}) {
	r.HTML(200, view, data)
}

// 获取区间内的随机数
func GetRandomByMinMax(min, max int) int{

	if max == min{

		return min

	}

	temp := max - min

	return rand.Intn(temp) + min

}

func IsExists(array []int, key int) bool{

	for _,items := range array{
		if items == key{

			return true

		}

	}
	return false
}

func IsExistsString(array []string, key string) bool{

	for _,items := range array{
		if items == key{

			return true

		}

	}
	return false
}
