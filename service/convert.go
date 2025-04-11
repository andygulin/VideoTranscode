package service

import (
	"bufio"
	"fmt"
	"os/exec"
)

type Converter interface {
	Convert()
}

type ConversionConfig struct {
	InputFile  string
	OutputFile string
}

type VideoCommandExecutor struct {
}

func (command *VideoCommandExecutor) ExecuteCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	fmt.Println(cmd.String())

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	stderrScanner := bufio.NewScanner(stderr)
	stdoutScanner := bufio.NewScanner(stdout)

	go func() {
		for stderrScanner.Scan() {
			fmt.Println(stderrScanner.Text())
		}
	}()

	go func() {
		for stdoutScanner.Scan() {
			fmt.Println(stdoutScanner.Text())
		}
	}()

	return cmd.Wait()
}
