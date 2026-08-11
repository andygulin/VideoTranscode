package process

import (
	"VideoTranscode/service"
	"fmt"
)

// VideoScale 视频缩放
type VideoScale struct {
	service.ConversionConfig
	service.VideoCommandExecutor

	Height int
	Width  int
}

func (obj *VideoScale) Convert() error {
	arg := []string{"-i", obj.InputFile}
	arg = append(arg, "-filter:v")
	arg = append(arg, fmt.Sprintf("scale=%d:%d", obj.Width, obj.Height))
	arg = append(arg, "-c:a")
	arg = append(arg, "copy")
	arg = append(arg, obj.OutputFile)

	err := obj.ExecuteCommand(service.GetMName(), arg...)
	if err != nil {
		return fmt.Errorf("failed to convert video: %w", err)
	}
	return nil
}
