package process

import (
	. "VideoTranscode/service"
)

// VideoAudioExtractor 提取视频中的音频
type VideoAudioExtractor struct {
	ConversionConfig
	VideoCommandExecutor
}

func (obj *VideoAudioExtractor) Convert() {
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

	err := obj.ExecuteCommand(GetMName(), arg...)
	if err != nil {
		panic(err)
	}
}
