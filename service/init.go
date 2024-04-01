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

	pathExists := func(path string) (bool, error) {
		_, err := os.Stat(path)
		if err == nil {
			return true, nil
		}
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	mName, _ = filepath.Abs(filepath.Join(pwd, "ffmpeg", osName, mCmd))
	exist, err := pathExists(mName)
	if !exist || err != nil {
		fmt.Printf("File Not Found : %s\n", mName)
	}

	pName, _ = filepath.Abs(filepath.Join(pwd, "ffmpeg", osName, pCmd))
	exist, err = pathExists(pName)
	if !exist || err != nil {
		fmt.Printf("File Not Found : %s\n", pName)
	}
}

func GetMName() string {
	return mName
}

func GetPName() string {
	return pName
}
