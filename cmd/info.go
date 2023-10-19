package cmd

import (
	. "VideoTranscode/service"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"time"
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
		obj := Info{}
		info, err := obj.GetInfo(args[0])
		if err != nil {
			fmt.Printf("Error : %s\n", err.Error())
			return
		}
		fmt.Printf("FileName: %s\n", info.Format.Filename)

		duration := info.Format.Duration
		f, _ := strconv.ParseFloat(duration, 64)
		fmt.Printf("Duration: %s\n", formatTime(int64(f)))

		size := info.Format.Size
		i, _ := strconv.ParseInt(size, 10, 64)
		fmt.Printf("FileSize: %s\n", formatFileSize(i))

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

func formatFileSize(fileSize int64) string {
	if fileSize < 1024 {
		return fmt.Sprintf("%.2f B", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2f KiB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f MiB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f GiB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f TiB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2f PiB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

func formatTime(sec int64) string {
	duration := time.Duration(sec) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
