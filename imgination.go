package main

import (
	"flag"

	"github.com/senorprogrammer/imgination/modules"
)

func main() {
	dirPath := flag.String("dir", "./", "Path to image directory")
	feature := flag.String("func", "dup", "Which function to execute. Options: dim, dup, gps")
	minWidth := flag.Int("width", 640, "Minimum image width")
	minHeight := flag.Int("height", 480, "Minimum image height")
	flag.Parse()

	switch *feature {
	case "dim":
		modules.FindMinimumDimensions(dirPath, minWidth, minHeight)
	case "dup":
		modules.FindDuplicates(dirPath)
	case "gps":
		modules.FindGps(dirPath)
	}
}
