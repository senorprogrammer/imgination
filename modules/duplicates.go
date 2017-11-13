package modules

import (
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/logrusorgru/aurora"
	"gopkg.in/h2non/filetype.v1"
	"path/filepath"
)

/* -------------------- Public -------------------- */

func FindDuplicates(dirPath *string) {
	fmt.Printf("Scanning for duplicates in %s...\n", *dirPath)

	hashMap := make(map[string]CollisionTable)

	filepath.Walk(*dirPath, func(path string, f os.FileInfo, err error) error {
		if isImage(path) == true {
			imgFile := NewImageFile(path)

			if isCollision(hashMap, imgFile) == true {
				collTable := hashMap[imgFile.Hash]
				collTable.Append(imgFile)
				hashMap[imgFile.Hash] = collTable

				fmt.Print(Red("D"))
			} else {
				hashMap[imgFile.Hash] = NewCollisionTable(imgFile)

				fmt.Print(Green("*"))
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
		count = count + collTable.CollisionCount()
	}

	return count
}

func isCollision(hashMap map[string]CollisionTable, imgFile *ImageFile) bool {
	if _, ok := hashMap[imgFile.Hash]; ok {
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
	fmt.Println("\n")
	fmt.Printf("Found %d duplicates\n\n", collisionCount(hashMap))

	for _, collTable := range hashMap {
		if collTable.HasCollisions() {
			fmt.Println(collTable.Paths())
		}
	}
}
