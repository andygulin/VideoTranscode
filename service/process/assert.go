package process

import "VideoTranscode/service"

var (
	_ service.Converter = (*ExtractAudio)(nil)
	_ service.Converter = (*Snapshot)(nil)
	_ service.Converter = (*TsMerge)(nil)
	_ service.Converter = (*TsList)(nil)
	_ service.Converter = (*VideoScale)(nil)
	_ service.Converter = (*VideoCrop)(nil)
	_ service.Converter = (*VideoConverter)(nil)
)
