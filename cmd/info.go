package cmd

import (
	"VideoTranscode/service"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Output video/audio file information.",
	Long:  "Output video/audio file information.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Missing video/audio files\n.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		obj := service.Transcode{}
		info, err := obj.Info(args[0])
		if err != nil {
			fmt.Printf("Error : %s\n", err.Error())
			return
		}
		fmt.Printf("FileName: %s\n", info.Format.Filename)

		duration := info.Format.Duration
		f, _ := strconv.ParseFloat(duration, 64)
		fmt.Printf("Duration: %s\n", fmt.Sprintf("%.2f", f))

		size := info.Format.Size
		i, _ := strconv.ParseInt(size, 10, 64)
		fmt.Printf("FileSize: %s\n", byteCountString(i))

		for i, stream := range info.Streams {
			idx := i + 1
			fmt.Printf("Stream%d\n", idx)
			fmt.Printf("\tIndex: %d\n", stream.Index)
			fmt.Printf("\tCodecName: %s\n", stream.CodecName)
			fmt.Printf("\tCodecLongName: %s\n", stream.CodecLongName)
			fmt.Printf("\tProfile: %s\n", stream.Profile)
			fmt.Printf("\tCodecType: %s\n", stream.CodecType)
			if stream.CodecType == "video" {
				fmt.Printf("\tSize: %d * %d\n", stream.Width, stream.Height)
				fmt.Printf("\tRatio: %s\n", stream.DisplayAspectRatio)
			}
		}
	},
}

func byteCountString(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
