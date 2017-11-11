package modules

import (
	"encoding/hex"
	"fmt"
	"image"
	"io/ioutil"
	"os"

	"github.com/devedge/imagehash"
	"github.com/gosuri/uitable"
	"gopkg.in/h2non/filetype.v1"
	"path/filepath"
)

/* -------------------- Public -------------------- */

func FindDuplicates(dirPath *string) {
	table := uitable.New()
	table.MaxColWidth = 80

	hashMap := make(map[string]string)

	collisionCount := 0
	filepath.Walk(*dirPath, func(path string, f os.FileInfo, err error) error {
		isImg := isImage(path)

		if isImg == true {
			hash, _ := hashFile(path)

			if checkForCollision(hashMap, hash) == true {
				table.AddRow(path, hashMap[hash])
				collisionCount += 1
			}

			hashMap[hash] = path
		}

		return nil
	})

	fmt.Printf("Found %d duplicates\n\n", collisionCount)
	fmt.Println(table)
}

/* -------------------- Private -------------------- */

func checkForCollision(hashMap map[string]string, hash string) bool {
	if _, ok := hashMap[hash]; ok {
		return true
	}
	return false
}

func hashFile(path string) (string, image.Image) {
	image, _ := imagehash.OpenImg(path)
	hash, _ := imagehash.Ahash(image, 16)

	return hex.EncodeToString(hash), image
}

func isImage(path string) bool {
	buf, _ := ioutil.ReadFile(path)

	if filetype.IsImage(buf) {
		return true
	}
	return false
}
