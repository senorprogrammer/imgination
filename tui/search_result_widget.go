package tui

import (
	"fmt"
	"os"
	"os/exec"
	_ "strconv"

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
		if err := g.SetKeybinding(widget.name, gocui.KeyArrowUp, gocui.ModNone, widget.selectPrev); err != nil {
			panic(err)
		}

		if err := g.SetKeybinding(widget.name, gocui.KeyArrowDown, gocui.ModNone, widget.selectNext); err != nil {
			panic(err)
		}

		if err := g.SetKeybinding(widget.name, gocui.KeyEnter, gocui.ModNone, widget.openFile); err != nil {
			panic(err)
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
				return err
			}
		}
	}
	return nil
}

func (widget *SearchResultWidget) cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		/* Returns the current position of the cursor */
		cx, cy := v.Cursor()

		/* Can't move down any further; don't go off the bottom of the list */
		if cy >= len(widget.body)-1 {
			return nil
		}

		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func (widget *SearchResultWidget) openFile(g *gocui.Gui, v *gocui.View) error {
	path := widget.body[widget.selected]
	if err := exec.Command("open", path).Run(); err != nil {
		return err
	}
	return nil
}

/* Tells the info view to display info about the selected file */
func (widget *SearchResultWidget) selectFile(g *gocui.Gui, v *gocui.View) error {
	path := widget.body[widget.selected]
	widget.handler(g, path)
	return nil
}

func (widget *SearchResultWidget) selectNext(g *gocui.Gui, v *gocui.View) error {
	err := widget.cursorDown(g, v)
	if err == nil && widget.selected < len(widget.body)-1 {
		widget.selected = widget.selected + 1
		widget.selectFile(g, v)
	}
	return err
}

func (widget *SearchResultWidget) selectPrev(g *gocui.Gui, v *gocui.View) error {
	err := widget.cursorUp(g, v)
	if err == nil && widget.selected != 0 {
		widget.selected = widget.selected - 1
		widget.selectFile(g, v)
	}
	return err
}
