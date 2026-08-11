package cmd

import (
	"VideoTranscode/service"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display application and FFmpeg version information",
	Long: `Print version information, including built‑in FFmpeg runtime version.

Example:
  VideoTranscode version
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return errors.New("version command does not accept any arguments")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := service.CheckFFTools(); err != nil {
			return err
		}
		obj := service.Version{}
		output, err := obj.GetVersion()
		if err != nil {
			return fmt.Errorf("failed to get version info: %w", err)
		}
		fmt.Println(output)
		return nil
	},
}
