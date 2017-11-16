package modules

import (
	"fmt"
	"os"
	"sync"

	. "github.com/logrusorgru/aurora"
	"github.com/stretchr/powerwalk"
)

func FindMinimumDimensions(dirPath *string, minWidth, minHeight *int) {
	fmt.Printf("Scanning %s for minimum dimensions (%d, %d)...\n", *dirPath, *minWidth, *minHeight)

	imgArray := []string{}
	var imgArrayLock sync.Mutex

	powerwalk.Walk(*dirPath, func(path string, f os.FileInfo, err error) error {
		if IsImage(path) == true {
			imgFile := NewImageFile(path)

			/*
			* Powerwalk scans files concurrently. Lock the storage array for each write
			 */
			imgArrayLock.Lock()
			defer imgArrayLock.Unlock()

			if imgFile.BelowMinimumDimensions(minWidth, minHeight) == true {
				imgArray = append(imgArray, imgFile.Path)

				fmt.Print(Red("S"))
			} else {
				fmt.Print(Green("*"))
			}
		}

		return nil
	})

	renderMinDimResults(imgArray)
}

func renderMinDimResults(imgArray []string) {
	fmt.Println("\n")
	fmt.Printf("Found %d images\n\n", len(imgArray))

	for _, path := range imgArray {
		fmt.Println(path)
	}
}
