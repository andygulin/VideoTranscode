package process

import (
	. "VideoTranscode/service"
	"bytes"
	"fmt"
	"os/exec"
)

// ConvertVideoMp3 提取视频中的音频
type ConvertVideoMp3 struct {
	Convert
}

func (obj *ConvertVideoMp3) Process() {
	arg := []string{"-i", obj.InputFile}
	arg = append(arg, "-vn")
	arg = append(arg, "-ar")
	arg = append(arg, "44100")
	arg = append(arg, "-ac")
	arg = append(arg, "2")
	arg = append(arg, "-ab")
	arg = append(arg, "320k")
	arg = append(arg, "-f")
	arg = append(arg, "mp3")
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
