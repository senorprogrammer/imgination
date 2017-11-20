package modules

import (
	"fmt"
	"os"
	"sync"

	. "github.com/logrusorgru/aurora"
	"github.com/senorprogrammer/imgination/image"
	"github.com/stretchr/powerwalk"
)

func FindDuplicates(dirPath *string) {
	fmt.Printf("Scanning %s for duplicates...\n", *dirPath)

	hashMap := make(map[string]CollisionTable)
	var lock sync.Mutex

	powerwalk.Walk(*dirPath, func(path string, f os.FileInfo, err error) error {
		if image.IsImage(path) == true {
			imgFile := image.NewImageFile(path)
			imgFile.GenerateHash()

			lock.Lock()
			defer lock.Unlock()

			if isCollision(hashMap, imgFile) == true {
				collTable := hashMap[imgFile.Hash]
				collTable.Append(imgFile)
				hashMap[imgFile.Hash] = collTable

				fmt.Print(Red("*"))
			} else {
				hashMap[imgFile.Hash] = NewCollisionTable(imgFile)

				fmt.Print(Green("*"))
			}
		}

		return nil
	})

	searchResult := SearchResult{}
	for _, collTable := range hashMap {
		if collTable.HasCollisions() {
			searchResult.Append(collTable.Paths())
		}
	}

	RenderTui(&searchResult)
}

/* -------------------- Private -------------------- */

func collisionCount(hashMap map[string]CollisionTable) int {
	count := 0
	for _, collTable := range hashMap {
		count = count + collTable.CollisionCount()
	}

	return count
}

func isCollision(hashMap map[string]CollisionTable, imgFile *image.ImageFile) bool {
	if _, ok := hashMap[imgFile.Hash]; ok {
		return true
	}
	return false
}
