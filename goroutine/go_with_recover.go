package goroutine

import (
	"log"
)

func Execute(f func()) {
	defer func() {
		cached := recover()
		if recovered != nil {
			log.Println(cached)
		}
	}()

	f()
}
