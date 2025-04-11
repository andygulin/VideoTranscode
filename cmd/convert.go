package cmd

import (
	. "VideoTranscode/service"
	. "VideoTranscode/service/process"
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
		obj := VideoConverter{
			ConversionConfig: ConversionConfig{
				InputFile:  args[0],
				OutputFile: args[1],
			},
		}
		if len(args) == 2 {
			obj.Convert()
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
			obj.Convert()
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
		obj := VideoAudioExtractor{
			ConversionConfig: ConversionConfig{
				InputFile:  args[0],
				OutputFile: args[1],
			},
		}
		obj.Convert()
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
		obj := VideoScale{
			ConversionConfig: ConversionConfig{
				InputFile:  args[0],
				OutputFile: args[1],
			},
			Width:  width,
			Height: height,
		}
		obj.Convert()
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
		obj := VideoFrameExtractor{
			ConversionConfig: ConversionConfig{
				InputFile:  args[0],
				OutputFile: "",
			},
		}
		obj.Convert()
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
		obj := VideoCrop{
			ConversionConfig: ConversionConfig{
				InputFile:  args[0],
				OutputFile: args[1],
			},
			StartTime: args[2],
			EndTime:   args[3],
		}
		obj.Convert()
	},
}

var ConvertGenerateTsListCmd = &cobra.Command{
	Use:   "generate_ts_list",
	Short: "The ts file list is generated.",
	Long:  "The ts file list is generated.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		obj := VideoTsListGenerator{
			ConversionConfig: ConversionConfig{
				InputFile:  args[0],
				OutputFile: args[1],
			},
		}
		obj.Convert()
	},
}

var ConvertMergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge video.",
	Long:  "Merge video.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		obj := VideoMerge{
			ConversionConfig: ConversionConfig{
				InputFile:  args[0],
				OutputFile: args[1],
			},
		}
		obj.Convert()
	},
}
