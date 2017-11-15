package modules

import (
	"fmt"
	"os"
	"sync"

	. "github.com/logrusorgru/aurora"
	"github.com/stretchr/powerwalk"
)

func FindGps(dirPath *string) {
	fmt.Printf("Scanning for locations in %s...\n", *dirPath)

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
