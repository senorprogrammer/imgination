package modules

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"github.com/gosuri/uitable"
	"gopkg.in/h2non/filetype.v1"
	"path/filepath"
)

/* -------------------- Public -------------------- */

type CollisionTable struct {
	Hash       string
	ImageFiles []*ImageFile
}

func (collTable *CollisionTable) Append(imgFile *ImageFile) {
	collTable.ImageFiles = append(collTable.ImageFiles, imgFile)
}

func (collTable *CollisionTable) HasCollisions() bool {
	return len(collTable.ImageFiles) > 1
}

/* -------------------- Main -------------------- */

func FindDuplicates(dirPath *string) {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	hashMap := make(map[string]CollisionTable)

	filepath.Walk(*dirPath, func(path string, f os.FileInfo, err error) error {
		if isImage(path) == true {
			imgFile := NewImageFile(path)

			if isCollision(hashMap, imgFile.Hash) == true {
				collTable := hashMap[imgFile.Hash]
				collTable.Append(imgFile)
				hashMap[imgFile.Hash] = collTable

				fmt.Printf("%s", red("D"))
			} else {
				hashMap[imgFile.Hash] = CollisionTable{
					Hash:       imgFile.Hash,
					ImageFiles: []*ImageFile{imgFile},
				}

				fmt.Printf("%s", green("*"))
			}
		}

		return nil
	})

	render(hashMap)

}

/* -------------------- Private -------------------- */

func collisionCount(hashMap map[string]CollisionTable) int {
	count := 0

	for _, collTable := range hashMap {
		if collTable.HasCollisions() {
			count++
		}
	}

	return count
}

func isCollision(hashMap map[string]CollisionTable, hash string) bool {
	if _, ok := hashMap[hash]; ok {
		return true
	}
	return false
}

func isImage(path string) bool {
	buf, _ := ioutil.ReadFile(path)

	if filetype.IsImage(buf) {
		return true
	}
	return false
}

func render(hashMap map[string]CollisionTable) {
	table := uitable.New()
	table.MaxColWidth = 80

	for _, collTable := range hashMap {
		if collTable.HasCollisions() {
			table.AddRow(collTable.Hash)
		}
	}

	fmt.Println("\n")
	fmt.Printf("Found %d duplicates\n\n", collisionCount(hashMap))
	fmt.Println(table)
}
