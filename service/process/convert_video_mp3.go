package process

import (
	. "VideoTranscode/service"
	"github.com/commander-cli/cmd"
	"strings"
)

// ConvertVideoMp3 提取视频中的音频
type ConvertVideoMp3 struct {
	Convert
}

func (obj *ConvertVideoMp3) Process() {
	var str []string
	str = append(str, GetMName())
	str = append(str, "-i")
	str = append(str, obj.InputFile)
	str = append(str, "-vn -ar 44100 -ac 2 -ab 320k -f mp3")
	str = append(str, obj.OutputFile)

	command := strings.Join(str, " ")
	c := cmd.NewCommand(command, cmd.WithStandardStreams)
	err := c.Execute()
	if err != nil {
		panic(err)
	}
}
