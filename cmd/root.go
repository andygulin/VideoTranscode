package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "VideoTranscode",
	Short: "Video/audio file processing.",
	Long:  "Video/audio file processing, including video transcoding, format conversion, video merging, screenshots...",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	rootCmd.AddCommand(VersionCmd, InfoCmd)

	ConvertCmd.AddCommand(
		ConvertVideoCmd,
		ConvertMp3Cmd,
		ConvertScaleCmd,
		ConvertImageCmd,
		ConvertCropCmd,
		ConvertMergeCmd,
		ConvertGenerateTsListCmd)
	rootCmd.AddCommand(ConvertCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
