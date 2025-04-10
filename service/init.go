package service

import (
	"fmt"
	"os/exec"
	"runtime"
)

var mName string
var pName string

func init() {
	mName = "ffmpeg"
	pName = "ffprobe"

	osName := runtime.GOOS
	if osName == "windows" {
		mName = mName + ".exe"
		pName = pName + ".exe"
	}

	checkCommandAvailable(mName)
	checkCommandAvailable(pName)
}

func checkCommandAvailable(command string) {
	_, err := exec.LookPath(command)
	if err != nil {
		panic(fmt.Sprintf("%s is not installed", command))
	}
}

func GetMName() string {
	return mName
}

func GetPName() string {
	return pName
}
