package modules

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/logrusorgru/aurora"
)

func FindGps(dirPath *string) {
	fmt.Printf("Scanning for locations in %s...\n", *dirPath)

	imgArray := []string{}

	filepath.Walk(*dirPath, func(path string, f os.FileInfo, err error) error {
		if IsImage(path) == true {
			imgFile := NewImageFile(path)

			if imgFile.HasGPS() == true {
				imgArray = append(imgArray, imgFile.Path)

				fmt.Print(Red("G"))
			} else {
				fmt.Print(Green("*"))
			}
		}

		return nil
	})

	renderGpsResults(imgArray)
}

func renderGpsResults(imgArray []string) {
	fmt.Println("\n")
	fmt.Printf("Found %d images\n\n", len(imgArray))

	for _, path := range imgArray {
		fmt.Println(path)
	}
}
