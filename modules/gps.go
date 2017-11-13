package modules

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	. "github.com/logrusorgru/aurora"
	"github.com/olekukonko/tablewriter"
)

func FindGps(dirPath *string) {
	fmt.Printf("Scanning for locations in %s...\n", *dirPath)

	imgArray := []*ImageFile{}

	filepath.Walk(*dirPath, func(path string, f os.FileInfo, err error) error {
		if IsImage(path) == true {
			imgFile := NewImageFile(path)

			if imgFile.HasGPS() == true {
				imgArray = append(imgArray, imgFile)

				fmt.Print(Red("G"))
			} else {
				fmt.Print(Green("*"))
			}
		}

		return nil
	})

	renderGpsResults(imgArray)
}

func renderGpsResults(imgArray []*ImageFile) {
	fmt.Println("\n")
	fmt.Printf("Found %d images\n\n", len(imgArray))

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Image", "Lat", "Lon"})

	for _, imgFile := range imgArray {
		lat, lon := imgFile.LatLon()

		data := []string{
			imgFile.Path,
			strconv.FormatFloat(lat, 'f', -1, 64),
			strconv.FormatFloat(lon, 'f', -1, 64),
		}
		table.Append(data)
	}

	table.Render()
}
