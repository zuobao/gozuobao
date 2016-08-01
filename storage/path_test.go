package storage

import (
	"testing"
	"time"
)

func Test_createpath(t *testing.T) {
	ChatFilepath(ChatFileTypeFile, 1, 2, 0, time.Now())

	//	t.Error(p)
}
