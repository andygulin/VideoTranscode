package service

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Version struct{}

func (obj *Version) GetVersion() string {
	cmd := exec.Command(GetMName(), "-version")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Sprintf("Error : %s\n", err.Error())
	}

	return out.String()
}
