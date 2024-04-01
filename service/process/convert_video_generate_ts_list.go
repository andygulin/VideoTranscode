package process

import (
	. "VideoTranscode/service"
	"fmt"
	"os"
	"sort"
	"strings"
)

// ConvertVideoGenerateTsList 生成TS列表文件
type ConvertVideoGenerateTsList struct {
	Convert
}

func (obj *ConvertVideoGenerateTsList) Process() {
	input := obj.InputFile
	files, err := os.ReadDir(input)
	if err != nil {
		panic(err)
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
			panic(err)
		}
	}
	fmt.Println("The " + obj.OutputFile + " file is successfully written.")
}
