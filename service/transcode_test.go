package service

import (
	"fmt"
	"testing"
)

var obj Transcode

func init() {
	obj = Transcode{}
}

func TestTranscode_Version(t *testing.T) {
	output := obj.Version()
	fmt.Println(output)
}
