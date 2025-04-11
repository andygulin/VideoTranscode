package process

import (
	. "VideoTranscode/service"
)

// VideoMerge ts合并mp4
type VideoMerge struct {
	ConversionConfig
	VideoCommandExecutor
}

func (obj *VideoMerge) Convert() {
	arg := []string{"-f", "concat", "-safe", "0", "-i"}
	arg = append(arg, obj.InputFile)
	arg = append(arg, "-c")
	arg = append(arg, "copy")
	arg = append(arg, obj.OutputFile)

	err := obj.ExecuteCommand(GetMName(), arg...)
	if err != nil {
		panic(err)
	}
}
