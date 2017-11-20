package modules

import (
	"encoding/hex"
	_ "fmt"
	"image"
	"io/ioutil"
	"os"

	"github.com/devedge/imagehash"
	"github.com/rwcarlsen/goexif/exif"
	"gopkg.in/h2non/filetype.v1"
)

type ImageFile struct {
	Hash   string
	Height int
	Image  image.Image
	Name   string
	Path   string
	Size   int64
	Width  int
}

func NewImageFile(path string) *ImageFile {
	imgFile := ImageFile{
		Hash:   "",
		Height: 0,
		Name:   "",
		Path:   path,
		Size:   0,
		Width:  0,
	}

	imgFile.loadFileDimensions()
	imgFile.loadFileInfo()

	return &imgFile
}

/* -------------------- Global Functions -------------------- */

func IsImage(path string) bool {
	buf, _ := ioutil.ReadFile(path)
	return filetype.IsImage(buf)
}

/* -------------------- Public Functions -------------------- */

func (imgFile *ImageFile) BelowMinimumDimensions(minWidth, minHeight *int) bool {
	if (imgFile.Width < *minWidth) || (imgFile.Height < *minHeight) {
		return true
	} else {
		return false
	}
}

func (imgFile *ImageFile) GenerateHash() {
	imgFile.loadImage()

	bytes, err := imagehash.Ahash(imgFile.Image, 16)
	if err == nil {
		imgFile.Hash = hex.EncodeToString(bytes)
	}

}

func (imgFile *ImageFile) HasGPS() bool {
	lat, lon := imgFile.LatLon()
	if lat != 0 && lon != 0 {
		return true
	}
	return false
}

func (imgFile *ImageFile) LatLon() (lat, lon float64) {
	file, err := os.Open(imgFile.Path)
	defer file.Close()

	if err != nil {
		return
	}

	ex, err := exif.Decode(file)
	if err != nil {
		return
	}

	lat, lon, _ = ex.LatLong()

	return lat, lon
}

/* -------------------- Private Functions -------------------- */

func (imgFile *ImageFile) loadFileDimensions() {
	file, err := os.Open(imgFile.Path)
	defer file.Close()

	if err != nil {
		return
	}

	conf, _, err := image.DecodeConfig(file)
	if err == nil {
		imgFile.Width, imgFile.Height = conf.Width, conf.Height
	}
}

func (imgFile *ImageFile) loadFileInfo() {
	info, err := os.Stat(imgFile.Path)
	if err == nil {
		imgFile.Name = info.Name()
		imgFile.Size = info.Size()
	}
}

/*
* This is a very expensive operation so it's deferred until absolutely needed
 */
func (imgFile *ImageFile) loadImage() {
	if imgFile.Image == nil {
		img, err := imagehash.OpenImg(imgFile.Path)
		if err != nil {
			return
		} else {
			imgFile.Image = img
		}
	}
}
