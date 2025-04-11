package process

import (
	. "VideoTranscode/service"
	"path/filepath"
	"strconv"
)

const DefaultSegmentTime = 10

// VideoConverter 视频文件格式转换
// mp4 -> avi
// mp4 -> mpeg
type VideoConverter struct {
	ConversionConfig
	VideoCommandExecutor

	// Lossless 无损转换
	Lossless bool

	// Segment 是否转m3u8
	Segment bool
	// SegmentTime m3u8 ts每个切片的时长（秒）
	SegmentTime int
}

func (obj *VideoConverter) Convert() {
	arg := []string{"-i"}
	arg = append(arg, obj.InputFile)
	if obj.Lossless {
		arg = append(arg, "-q:v")
		arg = append(arg, "0")
	}
	if obj.Segment {
		arg = append(arg, "-codec")
		arg = append(arg, "copy")
		arg = append(arg, "-bsf:v")
		arg = append(arg, "h264_mp4toannexb")
		arg = append(arg, "-map")
		arg = append(arg, "0")
		arg = append(arg, "-f")
		arg = append(arg, "segment")
		arg = append(arg, "-segment_list")
	}
	arg = append(arg, obj.OutputFile)
	if obj.Segment {
		if obj.SegmentTime > 0 {
			arg = append(arg, "-segment_time")
			arg = append(arg, strconv.Itoa(obj.SegmentTime))
		} else {
			arg = append(arg, "-segment_time")
			arg = append(arg, strconv.Itoa(DefaultSegmentTime))
		}
		arg = append(arg, filepath.Dir(obj.OutputFile)+string(filepath.Separator)+"%03d.ts")
	}

	err := obj.ExecuteCommand(GetMName(), arg...)
	if err != nil {
		panic(err)
	}
}
