package cmd

import (
	"VideoTranscode/service"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info <file> [file...]",
	Short: "Display metadata and stream information for video/audio files",
	Long: `Display detailed metadata, container and stream information for video or audio files.
It parses format metadata, video track parameters, audio track parameters, duration, bitrate, file size and other properties.

Examples:
  VideoTranscode info demo.mp4
  VideoTranscode info video.mp4 audio.aac
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("missing video/audio file argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := service.CheckFFTools(); err != nil {
			return err
		}
		obj := service.Info{}
		const splitLine = "========================================================"

		for _, filePath := range args {
			info, err := obj.GetInfo(filePath)
			if err != nil {
				fmt.Printf("\n[ERROR] get info for %s failed: %v\n", filePath, err)
				continue
			}

			fmt.Println("\n" + splitLine)
			fmt.Println("                     FILE INFORMATION")
			fmt.Println(splitLine)

			durationSec, _ := strconv.ParseFloat(info.Format.Duration, 64)
			sizeByte, _ := strconv.ParseInt(info.Format.Size, 10, 64)
			totalBitrate, _ := strconv.Atoi(info.Format.BitRate)

			fmt.Printf("File Path     : %s\n", info.Format.Filename)
			fmt.Printf("Container     : %s | %s\n", info.Format.FormatName, info.Format.FormatLongName)
			fmt.Printf("Duration      : %s (%.2f seconds)\n", formatDuration(int64(durationSec)), durationSec)
			fmt.Printf("File Size     : %s (%s bytes)\n", formatFileSize(sizeByte), info.Format.Size)
			fmt.Printf("Total Bitrate : %s\n", formatBitRate(totalBitrate))

			fmt.Println("\n" + splitLine)
			fmt.Println("                   STREAM DETAILS")
			fmt.Println(splitLine)

			for idx, stream := range info.Streams {
				streamNo := idx + 1
				var trackType string
				switch stream.CodecType {
				case "video":
					trackType = "[VIDEO]"
				case "audio":
					trackType = "[AUDIO]"
				default:
					trackType = fmt.Sprintf("[%s]", stream.CodecType)
				}

				fmt.Printf("\n▶ Stream #%d %s\n", streamNo, trackType)
				fmt.Println("--------------------------------------------------------")

				fmt.Printf("%-16s: %d\n", "Index", stream.Index)
				fmt.Printf("%-16s: %s\n", "Codec", stream.CodecName)
				fmt.Printf("%-16s: %s\n", "Codec Description", stream.CodecLongName)
				if stream.Profile != "" {
					fmt.Printf("%-16s: %s\n", "Profile", stream.Profile)
				}

				if stream.BitRate != "" {
					br, _ := strconv.Atoi(stream.BitRate)
					fmt.Printf("%-16s: %s\n", "Bitrate", formatBitRate(br))
				}

				if stream.NbFrames != "" {
					fmt.Printf("%-16s: %s frames\n", "Frames", stream.NbFrames)
				}

				if stream.CodecType == "video" {
					fmt.Printf("%-16s: %d × %d\n", "Resolution", stream.Width, stream.Height)
					fmt.Printf("%-16s: %s\n", "Display Aspect", stream.DisplayAspectRatio)
					fmt.Printf("%-16s: %s\n", "Pixel Format", stream.PixFmt)
					fmt.Printf("%-16s: %s\n", "Real FPS", stream.RFrameRate)
					fmt.Printf("%-16s: %s\n", "Avg FPS", stream.AvgFrameRate)
					if stream.Level > 0 {
						fmt.Printf("%-16s: %d\n", "Level", stream.Level)
					}
					if stream.Refs > 0 {
						fmt.Printf("%-16s: %d\n", "Ref Frames", stream.Refs)
					}
				}

				if stream.CodecType == "audio" {
					fmt.Printf("%-16s: %s Hz\n", "Sample Rate", stream.SampleRate)
					fmt.Printf("%-16s: %d\n", "Channels", stream.Channels)
					fmt.Printf("%-16s: %s\n", "Channel Layout", stream.ChannelLayout)
					fmt.Printf("%-16s: %s\n", "Sample Format", stream.SampleFmt)
					if stream.BitsPerSample > 0 {
						fmt.Printf("%-16s: %d bit\n", "Bit Depth", stream.BitsPerSample)
					}
				}

				if stream.Tags.Language != "" {
					fmt.Printf("%-16s: %s\n", "Language", stream.Tags.Language)
				}
			}
		}

		fmt.Println("\n" + splitLine)
		return nil
	},
}

func formatDuration(sec int64) string {
	if sec <= 0 {
		return "00:00:00"
	}
	dur := time.Duration(sec) * time.Second
	h := int(dur.Hours())
	m := int(dur.Minutes()) % 60
	s := int(dur.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func formatFileSize(fileSize int64) string {
	if fileSize < 1024 {
		return fmt.Sprintf("%.2f B", float64(fileSize))
	} else if fileSize < 1024*1024 {
		return fmt.Sprintf("%.2f KiB", float64(fileSize)/1024)
	} else if fileSize < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MiB", float64(fileSize)/1024/1024)
	} else {
		return fmt.Sprintf("%.2f GiB", float64(fileSize)/1024/1024/1024)
	}
}

func formatBitRate(bitRate int) string {
	if bitRate <= 0 {
		return "unknown"
	}
	if bitRate < 1000 {
		return fmt.Sprintf("%d bps", bitRate)
	} else if bitRate < 1000*1000 {
		return fmt.Sprintf("%.2f kbps", float64(bitRate)/1000)
	} else {
		return fmt.Sprintf("%.2f Mbps", float64(bitRate)/1000/1000)
	}
}
