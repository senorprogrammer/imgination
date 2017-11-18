package modules

import (
	"fmt"
	"log"
	_ "time"

	"github.com/jroimartin/gocui"
	"github.com/senorprogrammer/imgination/tui"
)

func Render(searchResult *SearchResult) {
	fmt.Println("\n")
	fmt.Printf("Found %d results\n\n", searchResult.Len())

	for _, result := range searchResult.Results {
		fmt.Println(result)
	}
}

/* -------------------- TUI -------------------- */

func RenderTui(searchResult *SearchResult) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	/* Build the widgets that define the interface */
	srcWidget := tui.NewSearchResultWidget("files", " Files ", searchResult.Results, displayFile)
	infoWidget := tui.NewInfoWidget("info", " Info ", "")
	conWidget := tui.NewConsoleWidget("console", " Options ")

	g.SetManager(srcWidget, infoWidget, conWidget)

	/* Add keybindings */
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	/* And start the main event loop */
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func displayFile(g *gocui.Gui, path string) {
	// view, err := g.View("info")
	// if err == nil {
	// 	// Tell the widget to update it's path value with path

	// } else {
	// 	panic(err)
	// }
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
