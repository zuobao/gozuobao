package util

import "bytes"




type Buffer struct {
	*bytes.Buffer
}

func (me *Buffer)  Close() error {
	return nil
}


func NewBuffer(buf []byte) *Buffer {
	ret := &Buffer {}
	ret.Buffer = bytes.NewBuffer(buf)
	return ret
}
