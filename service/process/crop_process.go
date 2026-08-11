package process

import (
	"VideoTranscode/service"
	"fmt"
)

// VideoCrop 视频剪切
type VideoCrop struct {
	service.ConversionConfig
	service.VideoCommandExecutor

	StartTime string
	EndTime   string
}

func (obj *VideoCrop) Convert() error {
	arg := []string{"-i"}
	arg = append(arg, obj.InputFile)
	arg = append(arg, "-ss")
	arg = append(arg, obj.StartTime)
	arg = append(arg, "-codec")
	arg = append(arg, "copy")
	arg = append(arg, "-to")
	arg = append(arg, obj.EndTime)
	arg = append(arg, obj.OutputFile)

	err := obj.ExecuteCommand(service.GetMName(), arg...)
	if err != nil {
		return fmt.Errorf("failed to crop video: %w", err)
	}
	return nil
}
