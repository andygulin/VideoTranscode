package process

import (
	"VideoTranscode/service"
	"fmt"
	"path/filepath"
	"strconv"
)

const DefaultSegmentTime = 10

// VideoConverter 视频文件格式转换
// mp4 -> avi
// mp4 -> mpeg
// mp4 -> m3u8(hls分片)
type VideoConverter struct {
	service.ConversionConfig
	service.VideoCommandExecutor

	// Lossless 无损转换
	Lossless bool

	// Segment 是否转m3u8
	Segment bool
	// SegmentTime m3u8 ts每个切片的时长（秒）
	SegmentTime int

	Width     int
	Height    int
	StartTime string
	Duration  string
}

func (obj *VideoConverter) Convert() error {
	arg := []string{"-i", obj.InputFile}

	if obj.Lossless && !obj.Segment {
		arg = append(arg, "-q:v", "0")
	}

	if obj.Segment {
		arg = append(arg, "-codec", "copy")

		bsf, err := getBSF(obj.InputFile)
		if err != nil {
			return fmt.Errorf("failed to get bsf: %w", err)
		}
		arg = append(arg, "-bsf:v", bsf)
		arg = append(arg, "-map", "0:v")
		arg = append(arg, "-map", "0:a")
		arg = append(arg, "-f", "segment")

		segTime := obj.SegmentTime
		if segTime <= 0 {
			segTime = DefaultSegmentTime
		}
		arg = append(arg, "-segment_time", strconv.Itoa(segTime))

		arg = append(arg, "-segment_list", obj.OutputFile)
		tsTemplate := filepath.Join(filepath.Dir(obj.OutputFile), "%03d.ts")
		arg = append(arg, tsTemplate)
	} else {
		arg = append(arg, obj.OutputFile)
	}

	err := obj.ExecuteCommand(service.GetMName(), arg...)
	if err != nil {
		return fmt.Errorf("failed to execute conversion command: %w", err)
	}
	return nil
}

func getBSF(inputFile string) (string, error) {
	obj := service.Info{}
	info, err := obj.GetInfo(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to get video info: %w", err)
	}
	for _, stream := range info.Streams {
		if stream.CodecType == "video" {
			switch stream.CodecName {
			case "h264":
				return "h264_mp4toannexb", nil
			case "hevc":
				return "hevc_mp4toannexb", nil
			default:
				return "", fmt.Errorf("lossless hls not support video codec: %s", stream.CodecName)
			}
		}
	}
	return "", fmt.Errorf("no video stream found in file: %s", inputFile)
}
