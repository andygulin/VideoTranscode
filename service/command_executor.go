package service

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"sync"
)

type Converter interface {
	Convert() error
}

type ConversionConfig struct {
	InputFile  string
	OutputFile string
}

type VideoCommandExecutor struct {
}

func (command *VideoCommandExecutor) ExecuteCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	fmt.Printf("[CMD] %s\n", cmd.String())

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("create stdout pipe failed: %w", err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("create stderr pipe failed: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start command failed: %w", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func(r io.Reader) {
		defer wg.Done()
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		_ = scanner.Err()
	}(stdoutPipe)

	go func(r io.Reader) {
		defer wg.Done()
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		_ = scanner.Err()
	}(stderrPipe)

	waitErr := cmd.Wait()
	wg.Wait()

	if waitErr != nil {
		return fmt.Errorf("command execute failed: %w", waitErr)
	}
	return nil
}
