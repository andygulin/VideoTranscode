package service

type Convert struct {
	InputFile  string
	OutputFile string
}

type ConvertProcess interface {
	Process()
}
