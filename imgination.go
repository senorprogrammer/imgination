package main

import (
	"flag"

	"github.com/senorprogrammer/imgination/modules"
)

func main() {
	dirPath := flag.String("dir", "", "Path to image directory")
	flag.Parse()

	modules.FindDuplicates(dirPath)
}
