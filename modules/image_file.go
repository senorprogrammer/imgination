package modules

import (
	"encoding/hex"
	"image"
	"io/ioutil"
	"os"

	"github.com/devedge/imagehash"
	"github.com/rwcarlsen/goexif/exif"
	"gopkg.in/h2non/filetype.v1"
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

	img, err := imagehash.OpenImg(imgFile.Path)
	if err == nil {
		imgFile.Image = img
	}

	bytes, err := imagehash.Ahash(imgFile.Image, 16)
	if err == nil {
		imgFile.Hash = hex.EncodeToString(bytes)
	}

	return &imgFile
}

/* -------------------- Public Functions -------------------- */

func (imgFile *ImageFile) HasGPS() bool {
	lat, lon := imgFile.LatLon()
	if lat != 0 && lon != 0 {
		return true
	}
	return false
}

func IsImage(path string) bool {
	buf, _ := ioutil.ReadFile(path)

	if filetype.IsImage(buf) {
		return true
	}
	return false
}

func (imgFile *ImageFile) LatLon() (lat, lon float64) {
	file, err := os.Open(imgFile.Path)
	if err != nil {
		return lat, lon
	}
	defer file.Close()

	ex, err := exif.Decode(file)
	if err != nil {
		return lat, lon
	}

	lat, lon, _ = ex.LatLong()

	return lat, lon
}
