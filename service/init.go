package service

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var mName string
var pName string

func init() {
	pwd, _ := os.Getwd()
	osName := runtime.GOOS
	mCmd := "ffmpeg"
	if osName == "windows" {
		mCmd = mCmd + ".exe"
	}

	pCmd := "ffprobe"
	if osName == "windows" {
		pCmd = pCmd + ".exe"
	}

	mName, _ = filepath.Abs(filepath.Join(pwd, "ffmpeg", osName, mCmd))
	exist, err := PathExists(mName)
	if !exist || err != nil {
		fmt.Printf("File Not Found : %s\n", mName)
	}

	pName, _ = filepath.Abs(filepath.Join(pwd, "ffmpeg", osName, pCmd))
	exist, err = PathExists(pName)
	if !exist || err != nil {
		fmt.Printf("File Not Found : %s\n", pName)
	}
}
