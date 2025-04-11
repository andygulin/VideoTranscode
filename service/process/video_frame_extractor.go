package process

import (
	. "VideoTranscode/service"
	"path/filepath"
)

// VideoFrameExtractor 视频帧转图片
type VideoFrameExtractor struct {
	ConversionConfig
	VideoCommandExecutor
}

func (obj *VideoFrameExtractor) Convert() {
	arg := []string{"-i"}
	arg = append(arg, obj.InputFile)
	arg = append(arg, "-r")
	arg = append(arg, "1")
	arg = append(arg, "-f")
	arg = append(arg, "image2")
	arg = append(arg, filepath.Dir(obj.InputFile)+string(filepath.Separator)+"image-%5d.png")

	err := obj.ExecuteCommand(GetMName(), arg...)
	if err != nil {
		panic(err)
	}
}
