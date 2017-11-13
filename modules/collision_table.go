package modules

import ()

type CollisionTable struct {
	Hash       string
	ImageFiles []*ImageFile
}

func NewCollisionTable(imgFile *ImageFile) CollisionTable {
	collTable := CollisionTable{
		Hash:       imgFile.Hash,
		ImageFiles: []*ImageFile{imgFile},
	}

	return collTable
}

func (collTable *CollisionTable) Append(imgFile *ImageFile) {
	collTable.ImageFiles = append(collTable.ImageFiles, imgFile)
}

func (collTable *CollisionTable) CollisionCount() int {
	return len(collTable.ImageFiles) - 1
}

func (collTable *CollisionTable) HasCollisions() bool {
	return len(collTable.ImageFiles) > 1
}

func (collTable *CollisionTable) Paths() string {
	pathStr := ""

	for i := 0; i < len(collTable.ImageFiles); i++ {
		imgFile := collTable.ImageFiles[i]
		pathStr = pathStr + imgFile.Path + "\n"
	}

	return pathStr
}
