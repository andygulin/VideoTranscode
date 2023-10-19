package service

import (
	"fmt"
	"github.com/commander-cli/cmd"
)

type Version struct{}

func (obj *Version) GetVersion() string {
	command := fmt.Sprintf("%s -version", mName)
	c := cmd.NewCommand(command)
	err := c.Execute()
	if err != nil {
		return fmt.Sprintf("Error : %s\n", err.Error())
	}
	return c.Stdout()
}
