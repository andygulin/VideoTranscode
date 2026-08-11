package service

import (
	"fmt"
	"os/exec"
	"runtime"
)

var (
	ffmpegBin  string
	ffprobeBin string
)

func init() {
	ffmpegBin = "ffmpeg"
	ffprobeBin = "ffprobe"

	if runtime.GOOS == "windows" {
		ffmpegBin += ".exe"
		ffprobeBin += ".exe"
	}
}

func CheckFFTools() error {
	if _, err := exec.LookPath(ffmpegBin); err != nil {
		return fmt.Errorf("dependency missing: %s not found in PATH", ffmpegBin)
	}
	if _, err := exec.LookPath(ffprobeBin); err != nil {
		return fmt.Errorf("dependency missing: %s not found in PATH", ffprobeBin)
	}
	return nil
}

func GetMName() string {
	return ffmpegBin
}

func GetPName() string {
	return ffprobeBin
}
