package process

import (
	"VideoTranscode/service"
	"fmt"
	"os"
	"sort"
	"strings"
)

// TsMerge ts合并mp4
type TsMerge struct {
	service.ConversionConfig
	service.VideoCommandExecutor
}

func (obj *TsMerge) Convert() error {
	arg := []string{"-f", "concat", "-safe", "0", "-i"}
	arg = append(arg, obj.InputFile)
	arg = append(arg, "-c")
	arg = append(arg, "copy")
	arg = append(arg, obj.OutputFile)

	err := obj.ExecuteCommand(service.GetMName(), arg...)
	if err != nil {
		return fmt.Errorf("failed to merge TS files: %w", err)
	}
	return nil
}

// TsList 生成TS列表文件
type TsList struct {
	service.ConversionConfig
}

func (obj *TsList) Convert() error {
	input := obj.InputFile
	files, err := os.ReadDir(input)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	var tsFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".ts") {
			tsFiles = append(tsFiles, input+"/"+file.Name())
		}
	}
	sort.Strings(tsFiles)

	output, err := os.Create(obj.OutputFile)
	defer func(output *os.File) {
		_ = output.Close()
	}(output)
	for _, fileName := range tsFiles {
		_, err := output.WriteString("file '" + fileName + "'\n")
		if err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}
	fmt.Println("The " + obj.OutputFile + " file is successfully written.")
	return nil
}
