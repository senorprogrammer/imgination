package tui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type SearchResultWidget struct {
	name     string
	title    string
	x, y     int
	w, h     int
	body     []string
	selected int
	handler  func(g *gocui.Gui, path string)
}

func NewSearchResultWidget(name string, title string, body []string, handler func(g *gocui.Gui, path string)) *SearchResultWidget {
	widget := SearchResultWidget{
		name:     name,
		title:    title,
		body:     body,
		selected: 0,
		handler:  handler,
	}

	return &widget
}

func (widget *SearchResultWidget) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	widget.x = 1
	widget.y = 0
	widget.w = int(0.75 * float32(maxX))
	widget.h = maxY - 4

	view, err := g.SetView(widget.name, widget.x, widget.y, widget.w, widget.h)
	if err != gocui.ErrUnknownView {
		return err
	} else {
		view.Title = fmt.Sprintf("%s (%d files) ", widget.title, len(widget.body))
		view.Highlight = true
		view.SelBgColor = gocui.ColorGreen
		view.SelFgColor = gocui.ColorBlack
		view.Wrap = false

		// Key bindings
		if err := g.SetKeybinding(widget.name, gocui.KeyArrowUp, gocui.ModNone, widget.cursorUp); err != nil {
			panic(err)
		}

		if err := g.SetKeybinding(widget.name, gocui.KeyArrowDown, gocui.ModNone, widget.cursorDown); err != nil {
			panic(err)
		}

		if err := g.SetKeybinding(widget.name, gocui.KeyEnter, gocui.ModNone, widget.selectFile); err != nil {
			return err
		}

		g.SetCurrentView(widget.name)

		// Render the file list
		for _, result := range widget.body {
			fmt.Fprintln(view, result)
		}
	}

	return nil
}

func (widget *SearchResultWidget) cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				widget.selected = oy - 1
				return err
			}
		}
	}
	return nil
}

func (widget *SearchResultWidget) cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()

		/* Don't go off the bottom of the list */
		if cy >= len(widget.body)-1 {
			return nil
		}

		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				widget.selected = oy + 1
				return err
			}
		}
	}
	return nil
}

/* Tells the info view to display info about the selected file */
func (widget *SearchResultWidget) selectFile(g *gocui.Gui, v *gocui.View) error {
	widget.handler(g, "cats")
	return nil
}
