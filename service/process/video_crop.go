package process

import (
	. "VideoTranscode/service"
)

// VideoCrop 视频剪切
type VideoCrop struct {
	ConversionConfig
	VideoCommandExecutor

	StartTime string
	EndTime   string
}

func (obj *VideoCrop) Convert() {
	arg := []string{"-i"}
	arg = append(arg, obj.InputFile)
	arg = append(arg, "-ss")
	arg = append(arg, obj.StartTime)
	arg = append(arg, "-codec")
	arg = append(arg, "copy")
	arg = append(arg, "-to")
	arg = append(arg, obj.EndTime)
	arg = append(arg, obj.OutputFile)

	err := obj.ExecuteCommand(GetMName(), arg...)
	if err != nil {
		panic(err)
	}
}
