package modules

import (
	"encoding/hex"
	"image"

	"github.com/devedge/imagehash"
)

type ImageFile struct {
	Hash  string
	Image image.Image
	Path  string
}

func NewImageFile(path string) *ImageFile {
	imgFile := ImageFile{
		Path: path,
	}

	img, _ := imagehash.OpenImg(imgFile.Path)
	imgFile.Image = img

	bytes, _ := imagehash.Ahash(imgFile.Image, 16)
	imgFile.Hash = hex.EncodeToString(bytes)

	return &imgFile
}
