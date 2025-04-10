package process

import (
	. "VideoTranscode/service"
	"bytes"
	"fmt"
	"os/exec"
)

// ConvertVideoCrop 视频剪切
type ConvertVideoCrop struct {
	Convert
	StartTime string
	EndTime   string
}

func (obj *ConvertVideoCrop) Process() {
	arg := []string{"-i"}
	arg = append(arg, obj.InputFile)
	arg = append(arg, "-ss")
	arg = append(arg, obj.StartTime)
	arg = append(arg, "-codec")
	arg = append(arg, "copy")
	arg = append(arg, "-to")
	arg = append(arg, obj.EndTime)
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
