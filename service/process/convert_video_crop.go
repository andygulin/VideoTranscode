package process

import (
	. "VideoTranscode/service"
	"github.com/commander-cli/cmd"
	"strings"
)

// ConvertVideoCrop 视频剪切
type ConvertVideoCrop struct {
	Convert
	StartTime string
	EndTime   string
}

func (obj *ConvertVideoCrop) Process() {
	var str []string
	str = append(str, GetMName())
	str = append(str, "-i")
	str = append(str, obj.InputFile)
	str = append(str, "-ss "+obj.StartTime)
	str = append(str, "-codec copy")
	str = append(str, "-to "+obj.EndTime)
	str = append(str, obj.OutputFile)

	command := strings.Join(str, " ")
	c := cmd.NewCommand(command, cmd.WithStandardStreams)
	err := c.Execute()
	if err != nil {
		panic(err)
	}
}
