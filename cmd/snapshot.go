package cmd

import (
	"VideoTranscode/service"
	"VideoTranscode/service/process"
	"errors"

	"github.com/spf13/cobra"
)

var SnapshotCmd = &cobra.Command{
	Use:   "snapshot <input> <output>",
	Short: "Capture one frame every second from video",
	Long: `Capture one frame every second from video file.
Output filename should use printf‑style format such as out_%04d.jpg.
Supported image formats: jpg, png, webp.

Example:
  VideoTranscode snapshot input.mp4 out_%04d.jpg
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("requires exactly 2 arguments: <input‑video> <output‑pattern>")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := service.CheckFFTools(); err != nil {
			return err
		}
		obj := process.Snapshot{
			ConversionConfig: service.ConversionConfig{
				InputFile:  args[0],
				OutputFile: args[1],
			},
		}
		return obj.Convert()
	},
}
