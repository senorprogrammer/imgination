package tui

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/jroimartin/gocui"
)

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

type InfoWidget struct {
	name      string
	title     string
	x, y      int
	w, h      int
	Path      string
	ImageStat *ImageStat
}

func NewInfoWidget(name, title, path string) *InfoWidget {
	widget := InfoWidget{
		name:      name,
		title:     title,
		Path:      path,
		ImageStat: &ImageStat{},
	}

	return &widget
}

func (widget *InfoWidget) DisplayFile(path string) {
	stat := NewImageStat(path)
	widget.ImageStat = stat
}

func (widget *InfoWidget) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	widget.x = int(0.75*float32(maxX)) + 1
	widget.y = 0
	widget.w = maxX - 1
	widget.h = maxY - 4

	// Displays the image stats widget
	statsView, _ := g.SetView("stats", widget.x, widget.y, widget.w, 6)
	statsView.Title = fmt.Sprintf(" %s ", widget.ImageStat.Name)

	statsView.Clear()
	fmt.Fprintln(statsView, "\n")
	fmt.Fprintf(statsView, "%8s: %8d\n", "Width", widget.ImageStat.Width)
	fmt.Fprintf(statsView, "%8s: %8d\n", "Height", widget.ImageStat.Height)
	fmt.Fprintf(statsView, "%8s: %8s\n", "Size", humanize.Bytes(uint64(widget.ImageStat.Size)))

	// Displays the image preview widget
	infoView, _ := g.SetView(widget.name, widget.x, widget.y+7, widget.w, widget.h)
	infoView.Title = " Info "
	infoView.Wrap = false

	infoView.Clear()

	return nil
}
