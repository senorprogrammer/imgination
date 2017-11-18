package tui

import (
	_ "fmt"

	"github.com/jroimartin/gocui"
)

type InfoWidget struct {
	name  string
	title string
	x, y  int
	w, h  int
	path  string
}

func NewInfoWidget(name, title, path string) *InfoWidget {
	widget := InfoWidget{
		name:  name,
		title: title,
		path:  path,
	}

	return &widget
}

func (widget *InfoWidget) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	widget.x = int(0.75*float32(maxX)) + 1
	widget.y = 0
	widget.w = maxX - 1
	widget.h = maxY - 4

	view, err := g.SetView(widget.name, widget.x, widget.y, widget.w, widget.h)
	if err != gocui.ErrUnknownView {
		return err
	} else {
		view.Title = " Info "
		view.Wrap = false
	}

	return nil
}
