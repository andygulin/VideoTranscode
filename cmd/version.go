package cmd

import (
	. "VideoTranscode/service"
	"fmt"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output the FFMPEG version.",
	Long:  "Output the FFMPEG version.",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		obj := Version{}
		output := obj.GetVersion()
		fmt.Println(output)
	},
}
