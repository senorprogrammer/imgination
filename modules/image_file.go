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

	img, _ := imagehash.OpenImg(imgFile.Path)
	imgFile.Image = img

	bytes, _ := imagehash.Ahash(imgFile.Image, 16)
	imgFile.Hash = hex.EncodeToString(bytes)

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

	x, err := exif.Decode(file)
	if err != nil {
		return lat, lon
	}

	lat, lon, _ = x.LatLong()

	return lat, lon
}
