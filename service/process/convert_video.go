package process

import (
	. "VideoTranscode/service"
	"fmt"
	"github.com/commander-cli/cmd"
	"path/filepath"
	"strconv"
	"strings"
)

// ConvertVideo 视频文件格式转换
// mp4 -> avi
// mp4 -> mpeg
type ConvertVideo struct {
	Convert
	// Lossless 无损转换
	Lossless bool

	Segment bool
	// SegmentTime mp4 -> m3u8 ts每个切片的时长（秒）
	SegmentTime int
}

func (obj *ConvertVideo) Process() {
	var str []string
	str = append(str, GetMName())
	str = append(str, "-i")
	str = append(str, obj.InputFile)
	if obj.Lossless {
		str = append(str, "-q:v 0")
	}
	if obj.Segment {
		str = append(str, "-codec copy -vbsf h264_mp4toannexb -map 0 -f segment -segment_list")
	}
	str = append(str, obj.OutputFile)
	if obj.Segment {
		if obj.SegmentTime > 0 {
			str = append(str, "-segment_time")
			str = append(str, strconv.Itoa(obj.SegmentTime))
		} else {
			str = append(str, "-segment_time 10")
		}
		str = append(str, filepath.Dir(obj.OutputFile)+string(filepath.Separator)+"%03d.ts")
	}

	command := strings.Join(str, " ")
	c := cmd.NewCommand(command, cmd.WithStandardStreams)
	fmt.Println(command)
	err := c.Execute()
	if err != nil {
		panic(err)
	}
}
