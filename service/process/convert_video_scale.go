package process

import (
	. "VideoTranscode/service"
	"github.com/commander-cli/cmd"
	"strconv"
	"strings"
)

// ConvertVideoScale 视频缩放
type ConvertVideoScale struct {
	Convert
	Height int
	Width  int
}

func (obj *ConvertVideoScale) Process() {
	var str []string
	str = append(str, GetMName())
	str = append(str, "-i")
	str = append(str, obj.InputFile)
	str = append(str, "-filter:v scale="+strconv.Itoa(obj.Width)+":"+strconv.Itoa(obj.Height)+" -c:a copy")
	str = append(str, obj.OutputFile)

	command := strings.Join(str, " ")
	c := cmd.NewCommand(command, cmd.WithStandardStreams)
	err := c.Execute()
	if err != nil {
		panic(err)
	}
}
