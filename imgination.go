package main

import (
	"flag"

	"github.com/senorprogrammer/imgination/modules"
)

func main() {
	dirPath := flag.String("dir", "./", "Path to image directory")
	feature := flag.String("func", "dup", "Which function to execute. Options: dup, gps")
	flag.Parse()

	switch *feature {
	case "dup":
		modules.FindDuplicates(dirPath)
	case "gps":
		modules.FindGps(dirPath)
	}
}
