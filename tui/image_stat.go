package tui

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// TODO: Find a way to unite this with ImageFile. No point in having two of these
type ImageStat struct {
	Width  int
	Height int
	Path   string
	Name   string
	Size   int64
}

func NewImageStat(path string) *ImageStat {
	imgStat := ImageStat{
		Path: path,
	}

	imgStat.loadDimensions()
	imgStat.loadFileInfo()

	return &imgStat
}

func (imgStat *ImageStat) loadDimensions() {
	file, _ := os.Open(imgStat.Path)
	defer file.Close()

	conf, _, err := image.DecodeConfig(file)
	if err == nil {
		imgStat.Height = conf.Height
		imgStat.Width = conf.Width
	}
}

func (imgStat *ImageStat) loadFileInfo() {
	info, err := os.Stat(imgStat.Path)
	if err == nil {
		imgStat.Name = info.Name()
		imgStat.Size = info.Size()
	}
}
