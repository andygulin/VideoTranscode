package process

import (
	. "VideoTranscode/service"
	"bytes"
	"fmt"
	"os/exec"
)

// ConvertVideoScale 视频缩放
type ConvertVideoScale struct {
	Convert
	Height int
	Width  int
}

func (obj *ConvertVideoScale) Process() {
	arg := []string{"-i", obj.InputFile}
	arg = append(arg, "-filter:v")
	arg = append(arg, fmt.Sprintf("scale=%d:%d", obj.Width, obj.Height))
	arg = append(arg, "-c:a")
	arg = append(arg, "copy")
	arg = append(arg, obj.OutputFile)
	cmd := exec.Command(GetMName(), arg...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	fmt.Println(cmd.String())
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
