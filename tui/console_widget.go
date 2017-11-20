package tui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type ConsoleWidget struct {
	name  string
	title string
	x, y  int
	w, h  int
}

func NewConsoleWidget(name, title string) *ConsoleWidget {
	widget := ConsoleWidget{
		name:  name,
		title: title,
	}

	return &widget
}

func (widget *ConsoleWidget) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	widget.x = 1
	widget.y = maxY - 3
	widget.w = maxX - 1
	widget.h = maxY - 1

	view, err := g.SetView(widget.name, widget.x, widget.y, widget.w, widget.h)
	if err != gocui.ErrUnknownView {
		return err
	} else {
		view.Title = fmt.Sprintf("%s", widget.title)
		view.SelBgColor = gocui.ColorGreen
		view.SelFgColor = gocui.ColorBlack
		view.Wrap = false

		fmt.Fprintln(view, "<DEL> Delete | <RET> Open | <Ctl-C> Quit ")
	}

	return nil
}
