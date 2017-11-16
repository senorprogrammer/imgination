package modules

import (
	"fmt"
	"os"
	"sync"

	. "github.com/logrusorgru/aurora"
	"github.com/stretchr/powerwalk"
)

/* -------------------- Public -------------------- */

func FindDuplicates(dirPath *string) {
	fmt.Printf("Scanning %s for duplicates...\n", *dirPath)

	hashMap := make(map[string]CollisionTable)
	var hashMapLock sync.Mutex

	powerwalk.Walk(*dirPath, func(path string, f os.FileInfo, err error) error {
		if IsImage(path) == true {
			imgFile := NewImageFile(path)
			imgFile.GenerateHash()

			/*
			* Powerwalk scans files concurrently. Lock the storage array for each write
			 */
			hashMapLock.Lock()
			defer hashMapLock.Unlock()

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

	renderDuplicationResults(hashMap)
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

func renderDuplicationResults(hashMap map[string]CollisionTable) {
	fmt.Println("\n")
	fmt.Printf("Found %d duplicates\n\n", collisionCount(hashMap))

	for _, collTable := range hashMap {
		if collTable.HasCollisions() {
			fmt.Println(collTable.Paths())
		}
	}
}
