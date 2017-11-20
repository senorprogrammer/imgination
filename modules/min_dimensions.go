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

	searchResult := SearchResult{}
	var lock sync.Mutex

	powerwalk.Walk(*dirPath, func(path string, f os.FileInfo, err error) error {
		if IsImage(path) == true {
			imgFile := NewImageFile(path)

			lock.Lock()
			defer lock.Unlock()

			if imgFile.BelowMinimumDimensions(minWidth, minHeight) == true {
				searchResult.Append(imgFile.Path)

				fmt.Print(Red("*"))
			} else {
				fmt.Print(Green("*"))
			}
		}

		return nil
	})

	RenderTui(&searchResult)
}
