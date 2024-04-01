package process

import (
	. "VideoTranscode/service"
	"github.com/commander-cli/cmd"
	"strings"
)

// ConvertVideoMerge ts合并mp4
type ConvertVideoMerge struct {
	Convert
}

func (obj *ConvertVideoMerge) Process() {
	var str []string
	str = append(str, GetMName())
	str = append(str, "-f concat -safe 0")
	str = append(str, "-i")
	str = append(str, obj.InputFile)
	str = append(str, "-c copy")
	str = append(str, obj.OutputFile)
	command := strings.Join(str, " ")
	c := cmd.NewCommand(command, cmd.WithStandardStreams)
	err := c.Execute()
	if err != nil {
		panic(err)
	}
}
