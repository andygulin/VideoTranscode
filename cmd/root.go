package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "VideoTranscode",
	Short: "A powerful video and audio command‑line processing tool",
	Long: `VideoTranscode is a command‑line multimedia processing tool built on FFmpeg.
Provides capabilities including video transcoding, format conversion, audio extraction,
media information inspection, snapshot capture, TS stream merge and other functions.

Examples:
  VideoTranscode version
  VideoTranscode info input.mp4
  VideoTranscode transcode --lossless input.mp4 output.avi
  VideoTranscode extract‑audio input.mp4 output.mp3
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	rootCmd.AddCommand(VersionCmd)
	rootCmd.AddCommand(InfoCmd)
	rootCmd.AddCommand(ExtractAudioCmd)
	rootCmd.AddCommand(SnapshotCmd)
	rootCmd.AddCommand(TranscodeCmd)

	TsCmd.AddCommand(TsMergeCmd)
	TsCmd.AddCommand(TsListCmd)
	rootCmd.AddCommand(TsCmd)

	if err := rootCmd.Execute(); err != nil {
	}
}
