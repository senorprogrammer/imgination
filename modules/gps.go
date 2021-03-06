package modules

import (
	"fmt"
	"os"
	"sync"

	. "github.com/logrusorgru/aurora"
	"github.com/senorprogrammer/imgination/image"
	"github.com/stretchr/powerwalk"
)

func FindGps(dirPath *string) {
	fmt.Printf("Scanning %s for locations...\n", *dirPath)

	searchResult := SearchResult{}
	var lock sync.Mutex

	powerwalk.Walk(*dirPath, func(path string, f os.FileInfo, err error) error {
		if image.IsImage(path) == true {
			imgFile := image.NewImageFile(path)

			lock.Lock()
			defer lock.Unlock()

			if imgFile.HasGPS() == true {
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
