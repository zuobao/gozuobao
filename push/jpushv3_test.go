package push

import (
	"log"
	"testing"
)

func Test_jpushv3(t *testing.T) {
	jpushv3 := NewJPushV3Engine("", "", false)
	msgid, err := jpushv3.PushAll("", "", nil)
	log.Println(msgid, err)
}
