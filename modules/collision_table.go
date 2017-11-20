package modules

import (
	"github.com/senorprogrammer/imgination/image"
)

type CollisionTable struct {
	Hash       string
	ImageFiles []*image.ImageFile
}

func NewCollisionTable(imgFile *image.ImageFile) CollisionTable {
	collTable := CollisionTable{
		Hash:       imgFile.Hash,
		ImageFiles: []*image.ImageFile{imgFile},
	}

	return collTable
}

/* -------------------- Public Functions -------------------- */

func (collTable *CollisionTable) Append(imgFile *image.ImageFile) {
	collTable.ImageFiles = append(collTable.ImageFiles, imgFile)
}

func (collTable *CollisionTable) CollisionCount() int {
	return len(collTable.ImageFiles) - 1
}

func (collTable *CollisionTable) HasCollisions() bool {
	return collTable.CollisionCount() > 0
}

func (collTable *CollisionTable) Paths() string {
	pathStr := ""

	for i := 0; i < len(collTable.ImageFiles); i++ {
		imgFile := collTable.ImageFiles[i]
		pathStr = pathStr + imgFile.Path + "\n"
	}

	return pathStr
}
