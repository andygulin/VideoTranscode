package process

import (
	. "VideoTranscode/service"
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
)

// ConvertVideoImage 视频帧转图片
type ConvertVideoImage struct {
	Convert
}

func (obj *ConvertVideoImage) Process() {
	arg := []string{"-i"}
	arg = append(arg, obj.InputFile)
	arg = append(arg, "-r")
	arg = append(arg, "1")
	arg = append(arg, "-f")
	arg = append(arg, "image2")
	arg = append(arg, filepath.Dir(obj.InputFile)+string(filepath.Separator)+"image-%5d.png")
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
