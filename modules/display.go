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

type WidgetManagers struct {
	ConWidget          *tui.ConsoleWidget
	InfoWidget         *tui.InfoWidget
	SearchResultWidget *tui.SearchResultWidget
}

var widgetMan = WidgetManagers{}

func RenderTui(searchResult *SearchResult) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	/* Build the widgets that define the interface */
	widgetMan.ConWidget = tui.NewConsoleWidget("console", " Options ")
	widgetMan.InfoWidget = tui.NewInfoWidget("info", " Info ", "")
	widgetMan.SearchResultWidget = tui.NewSearchResultWidget("files", " Files ", searchResult.Results, displayFile)

	g.SetManager(widgetMan.ConWidget, widgetMan.InfoWidget, widgetMan.SearchResultWidget)

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
	widgetMan.InfoWidget.Path = path
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
