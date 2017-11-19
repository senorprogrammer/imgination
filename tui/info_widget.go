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
	Path  string
}

func NewInfoWidget(name, title, path string) *InfoWidget {
	widget := InfoWidget{
		name:  name,
		title: title,
		Path:  path,
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
		// I don't understand this, because the first time it's created it'll be this error type
		// Foreer after it'll be nil
		// fmt.Fprintln(view, err)
		// return err
	} else {
		view.Title = " Info "
		view.Wrap = false
	}

	// fmt.Fprintln(view, "cats")
	// fmt.Fprintln(view, widget.Path)

	return nil
}
