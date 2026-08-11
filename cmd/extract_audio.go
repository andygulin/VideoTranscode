package cmd

import (
	"VideoTranscode/service"
	"VideoTranscode/service/process"
	"errors"

	"github.com/spf13/cobra"
)

var ExtractAudioCmd = &cobra.Command{
	Use:   "extract-audio <input> <output>",
	Short: "Extract audio stream from video file",
	Long: `Extract audio stream from video file.
Supports extracting audio track from video and saving to audio formats such as mp3, aac, m4a.

Examples:
  VideoTranscode extract-audio input.mp4 output.mp3
  VideoTranscode extract-audio input.mp4 output.aac
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("requires exactly 2 arguments: <input-video> <output-audio>")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := service.CheckFFTools(); err != nil {
			return err
		}

		obj := &process.ExtractAudio{
			ConversionConfig: service.ConversionConfig{
				InputFile:  args[0],
				OutputFile: args[1],
			},
		}
		return obj.Convert()
	},
}
