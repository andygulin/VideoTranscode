package process

import (
	. "VideoTranscode/service"
	"bytes"
	"fmt"
	"os/exec"
)

// ConvertVideoMerge ts合并mp4
type ConvertVideoMerge struct {
	Convert
}

func (obj *ConvertVideoMerge) Process() {
	arg := []string{"-f", "concat", "-safe", "0", "-i"}
	arg = append(arg, obj.InputFile)
	arg = append(arg, "-c")
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
