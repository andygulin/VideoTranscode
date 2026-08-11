package process

import (
	"VideoTranscode/service"
	"fmt"
	"path/filepath"
)

// Snapshot 视频帧转图片
type Snapshot struct {
	service.ConversionConfig
	service.VideoCommandExecutor
}

func (obj *Snapshot) Convert() error {
	arg := []string{"-i"}
	arg = append(arg, obj.InputFile)
	arg = append(arg, "-r")
	arg = append(arg, "1")
	arg = append(arg, "-f")
	arg = append(arg, "image2")
	arg = append(arg, filepath.Dir(obj.InputFile)+string(filepath.Separator)+"image-%5d.png")

	err := obj.ExecuteCommand(service.GetMName(), arg...)
	if err != nil {
		return fmt.Errorf("failed to take snapshot: %w", err)
	}
	return nil
}
