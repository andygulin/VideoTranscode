package process

import (
	. "VideoTranscode/service"
	"github.com/commander-cli/cmd"
	"path/filepath"
	"strings"
)

// ConvertVideoImage 视频帧转图片
type ConvertVideoImage struct {
	Convert
}

func (obj *ConvertVideoImage) Process() {
	var str []string
	str = append(str, GetMName())
	str = append(str, "-i")
	str = append(str, obj.InputFile)
	str = append(str, "-r 1 -f image2")
	str = append(str, filepath.Dir(obj.InputFile)+string(filepath.Separator)+"image-%5d.png")

	command := strings.Join(str, " ")
	c := cmd.NewCommand(command, cmd.WithStandardStreams)
	err := c.Execute()
	if err != nil {
		panic(err)
	}
}
