package tui

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/jroimartin/gocui"
)

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
	widget.Path = path
	stat := NewImageStat(widget.Path)
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
	fmt.Fprint(infoView, InlineImage(infoView, widget.Path))
	// fmt.Print(InlineImage(infoView, widget.Path))

	return nil
}
