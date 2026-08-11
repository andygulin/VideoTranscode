package process

import (
	"VideoTranscode/service"
	"fmt"
)

// ExtractAudio 提取视频中的音频
type ExtractAudio struct {
	service.ConversionConfig
	service.VideoCommandExecutor
}

func (obj *ExtractAudio) Convert() error {
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

	err := obj.ExecuteCommand(service.GetMName(), arg...)
	if err != nil {
		return fmt.Errorf("failed to execute audio extraction command: %w", err)
	}
	return nil
}
