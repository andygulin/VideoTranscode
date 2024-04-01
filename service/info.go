package service

import (
	"encoding/json"
	"fmt"
	"github.com/commander-cli/cmd"
)

type Info struct{}

func (obj *Info) GetInfo(file string) (FileInfo, error) {
	command := fmt.Sprintf("%s -v quiet -show_format -show_streams -print_format json %s", GetPName(), file)
	c := cmd.NewCommand(command)
	err := c.Execute()
	if err != nil {
		return FileInfo{}, err
	}
	output := c.Stdout()
	var fileInfo FileInfo
	err = json.Unmarshal([]byte(output), &fileInfo)
	if err != nil {
		return FileInfo{}, err
	}
	return fileInfo, nil
}
