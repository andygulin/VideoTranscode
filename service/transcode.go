package service

import (
	"encoding/json"
	"fmt"
	"github.com/commander-cli/cmd"
	"os"
	"path/filepath"
	"runtime"
)

var ffmpegName string
var ffprobeName string

func init() {
	pwd, _ := os.Getwd()
	osName := runtime.GOOS
	ffmpegCmd := "ffmpeg"
	if osName == "windows" {
		ffmpegCmd = ffmpegCmd + ".exe"
	}

	ffprobeCmd := "ffprobe"
	if osName == "windows" {
		ffprobeCmd = ffprobeCmd + ".exe"
	}

	ffmpegName, _ = filepath.Abs(filepath.Join(pwd, "ffmpeg", osName, ffmpegCmd))
	exist, err := PathExists(ffmpegName)
	if !exist || err != nil {
		fmt.Printf("File Not Found : %s\n", ffmpegName)
	}

	ffprobeName, _ = filepath.Abs(filepath.Join(pwd, "ffmpeg", osName, ffprobeCmd))
	exist, err = PathExists(ffprobeName)
	if !exist || err != nil {
		fmt.Printf("File Not Found : %s\n", ffprobeName)
	}
}

type Transcode struct{}

func (obj *Transcode) Version() string {
	command := fmt.Sprintf("%s -version", ffmpegName)
	c := cmd.NewCommand(command)
	err := c.Execute()
	if err != nil {
		return fmt.Sprintf("Error : %s\n", err.Error())
	}
	return c.Stdout()
}

func (obj *Transcode) Info(file string) (Info, error) {
	command := fmt.Sprintf("%s -v quiet -show_format -show_streams -print_format json %s", ffprobeName, file)
	c := cmd.NewCommand(command)
	err := c.Execute()
	if err != nil {
		return Info{}, err
	}
	output := c.Stdout()
	var info Info
	err = json.Unmarshal([]byte(output), &info)
	if err != nil {
		return Info{}, err
	}
	return info, nil
}
