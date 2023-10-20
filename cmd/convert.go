package cmd

import (
	. "VideoTranscode/service"
	"github.com/spf13/cobra"
	"path/filepath"
	"strconv"
)

var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Video transcoding.",
	Long:  "Video transcoding.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var ConvertVideoCmd = &cobra.Command{
	Use:   "video",
	Short: "Video format conversion.",
	Long:  "Video format conversion.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		obj := ConvertVideo{
			Convert: Convert{
				InputFile:  args[0],
				OutputFile: args[1],
			},
		}
		if len(args) == 2 {
			obj.Process()
		}
		if len(args) == 3 {
			outputFile := obj.OutputFile
			ext := filepath.Ext(outputFile)
			if ext == ".m3u8" {
				segmentTime, _ := strconv.Atoi(args[2])
				obj.Segment = true
				obj.SegmentTime = segmentTime
				obj.Lossless = true
			} else {
				lossless, _ := strconv.ParseBool(args[2])
				obj.Lossless = lossless
			}
			obj.Process()
		}
	},
}

var ConvertMp3Cmd = &cobra.Command{
	Use:   "mp3",
	Short: "Extract the audio from the video.",
	Long:  "Extract the audio from the video.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		obj := ConvertVideoMp3{
			Convert: Convert{
				InputFile:  args[0],
				OutputFile: args[1],
			},
		}
		obj.Process()
	},
}

var ConvertScaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Video scaling.",
	Long:  "Video scaling.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		width, _ := strconv.Atoi(args[2])
		height, _ := strconv.Atoi(args[3])
		obj := ConvertVideoScale{
			Convert: Convert{
				InputFile:  args[0],
				OutputFile: args[1],
			},
			Width:  width,
			Height: height,
		}
		obj.Process()
	},
}

var ConvertImageCmd = &cobra.Command{
	Use:   "image",
	Short: "Video to picture.",
	Long:  "Video to picture.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		obj := ConvertVideoImage{
			Convert: Convert{
				InputFile:  args[0],
				OutputFile: "",
			},
		}
		obj.Process()
	},
}

var ConvertCropCmd = &cobra.Command{
	Use:   "crop",
	Short: "Crop video.",
	Long:  "Crop video.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		obj := ConvertVideoCrop{
			Convert: Convert{
				InputFile:  args[0],
				OutputFile: args[1],
			},
			StartTime: args[2],
			EndTime:   args[3],
		}
		obj.Process()
	},
}
