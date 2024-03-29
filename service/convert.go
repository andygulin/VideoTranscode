package service

import (
	"fmt"
	"github.com/commander-cli/cmd"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Convert struct {
	InputFile  string
	OutputFile string
}

type ConvertProcess interface {
	Process()
}

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
	str = append(str, mName)
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

// ConvertVideoMp3 提取视频中的音频
type ConvertVideoMp3 struct {
	Convert
}

func (obj *ConvertVideoMp3) Process() {
	var str []string
	str = append(str, mName)
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

// ConvertVideoScale 视频缩放
type ConvertVideoScale struct {
	Convert
	Height int
	Width  int
}

func (obj *ConvertVideoScale) Process() {
	var str []string
	str = append(str, mName)
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

// ConvertVideoImage 视频帧转图片
type ConvertVideoImage struct {
	Convert
}

func (obj *ConvertVideoImage) Process() {
	var str []string
	str = append(str, mName)
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

// ConvertVideoCrop 视频剪切
type ConvertVideoCrop struct {
	Convert
	StartTime string
	EndTime   string
}

func (obj *ConvertVideoCrop) Process() {
	var str []string
	str = append(str, mName)
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

// ConvertVideoGenerateTsList 生成TS列表文件
type ConvertVideoGenerateTsList struct {
	Convert
}

func (obj *ConvertVideoGenerateTsList) Process() {
	input := obj.InputFile
	files, err := os.ReadDir(input)
	if err != nil {
		panic(err)
	}

	var tsFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".ts") {
			tsFiles = append(tsFiles, input+"/"+file.Name())
		}
	}
	sort.Strings(tsFiles)

	output, err := os.Create(obj.OutputFile)
	defer func(output *os.File) {
		_ = output.Close()
	}(output)
	for _, fileName := range tsFiles {
		_, err := output.WriteString("file '" + fileName + "'\n")
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("The " + obj.OutputFile + " file is successfully written.")
}

// ConvertVideoMerge ts合并mp4
type ConvertVideoMerge struct {
	Convert
}

func (obj *ConvertVideoMerge) Process() {
	var str []string
	str = append(str, mName)
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
