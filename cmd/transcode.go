package cmd

import (
	"VideoTranscode/service"
	"VideoTranscode/service/process"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	lossless bool
	hlsTime  int
	scaleStr string
	cropStr  string
)

var TranscodeCmd = &cobra.Command{
	Use:   "transcode [OPTIONS] <input> <output>",
	Short: "General video transcode",
	Long: `General video/audio transcode.

Examples:
  VideoTranscode transcode --lossless input.mp4 out.avi
  VideoTranscode transcode --lossless --hls-time 20 input.mp4 out.m3u8
  VideoTranscode transcode --scale 1280:720 input.mp4 out.mp4
  VideoTranscode transcode --crop start=00:00:00,duration=00:01:00 input.mp4 out.mp4
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("requires exactly 2 positional arguments: <input> <output>")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := service.CheckFFTools(); err != nil {
			return err
		}
		inputFile := args[0]
		outputFile := args[1]

		obj := process.VideoConverter{
			ConversionConfig: service.ConversionConfig{
				InputFile:  inputFile,
				OutputFile: outputFile,
			},
			Lossless: lossless,
		}

		ext := strings.ToLower(filepath.Ext(outputFile))
		if ext == ".m3u8" {
			obj.Segment = true
		}

		if hlsTime > 0 {
			obj.Segment = true
			obj.SegmentTime = hlsTime
		}

		if scaleStr != "" {
			parts := strings.Split(scaleStr, ":")
			if len(parts) != 2 {
				return fmt.Errorf("scale format error, expect W:H, e.g 1280:720, got %s", scaleStr)
			}
			w, err := strconv.Atoi(strings.TrimSpace(parts[0]))
			if err != nil {
				return fmt.Errorf("scale width invalid: %w", err)
			}
			h, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				return fmt.Errorf("scale height invalid: %w", err)
			}

			obj.Width = w
			obj.Height = h
		}

		if cropStr != "" {
			cropMap := make(map[string]string)
			items := strings.Split(cropStr, ",")
			for _, item := range items {
				kv := strings.SplitN(item, "=", 2)
				if len(kv) != 2 {
					return fmt.Errorf("crop param invalid: %s, expect key=value", item)
				}
				k := strings.TrimSpace(kv[0])
				v := strings.TrimSpace(kv[1])
				cropMap[k] = v
			}

			startVal, okStart := cropMap["start"]
			durVal, okDur := cropMap["duration"]
			if !okStart || !okDur {
				return errors.New("crop need start=xx,duration=xx, example: start=00:00:00,duration=00:01:00")
			}

			obj.StartTime = startVal
			obj.Duration = durVal
		}

		return obj.Convert()
	},
}

func init() {
	TranscodeCmd.Flags().BoolVarP(&lossless, "lossless", "l", false, "lossless stream copy")
	TranscodeCmd.Flags().IntVarP(&hlsTime, "hls-time", "", 0, "HLS segment time in seconds for m3u8 output")
	TranscodeCmd.Flags().StringVarP(&scaleStr, "scale", "", "", "scale video, format W:H e.g 1280:720")
	TranscodeCmd.Flags().StringVarP(&cropStr, "crop", "", "", "crop video, format start=xx,duration=xx")
}
