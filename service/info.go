package service

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

type Info struct{}

func (obj *Info) GetInfo(file string) (FileInfo, error) {
	arg := []string{"-v", "quiet", "-show_format", "-show_streams", "-print_format", "json", file}
	cmd := exec.Command(GetPName(), arg...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return FileInfo{}, err
	}
	output := out.String()

	var fileInfo FileInfo
	err = json.Unmarshal([]byte(output), &fileInfo)
	if err != nil {
		return FileInfo{}, err
	}
	return fileInfo, nil
}
