package util

import (
	"io"
	"io/ioutil"
	"encoding/json"
)




func ReadJson(reader io.ReadCloser, data interface {}) error {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	if bytes == nil {
		return nil
	}

	return json.Unmarshal(bytes, data)
}
