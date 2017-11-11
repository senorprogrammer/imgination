package modules

import (
	"encoding/hex"
	"fmt"
	"image"
	"io/ioutil"
	"os"

	"github.com/devedge/imagehash"
	"github.com/fatih/color"
	"github.com/gosuri/uitable"
	"gopkg.in/h2non/filetype.v1"
	"path/filepath"
)

/* -------------------- Public -------------------- */

func FindDuplicates(dirPath *string) {
	table := uitable.New()
	table.MaxColWidth = 80

	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	hashMap := make(map[string]string)

	collisionCount := 0
	filepath.Walk(*dirPath, func(path string, f os.FileInfo, err error) error {
		if isImage(path) == true {
			hash, _ := hashFile(path)

			if checkForCollision(hashMap, hash) == true {
				table.AddRow(path, hashMap[hash])
				collisionCount += 1

				fmt.Printf("%s", red("D"))
			} else {
				fmt.Printf("%s", green("*"))
			}

			hashMap[hash] = path
		}

		return nil
	})

	fmt.Println("\n")
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
