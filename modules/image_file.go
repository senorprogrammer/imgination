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
	Hash  string
	Image image.Image
	Path  string
}

func NewImageFile(path string) *ImageFile {
	imgFile := ImageFile{
		Path: path,
	}

	/*
	* TODO: Figure out why Dimensions doesn't work when this is commented out
	* This _should_ have no impact on that
	 */
	img, err := imagehash.OpenImg(imgFile.Path)
	if err == nil {
		imgFile.Image = img
	}

	return &imgFile
}

/* -------------------- Global Functions -------------------- */

func IsImage(path string) bool {
	buf, _ := ioutil.ReadFile(path)
	return filetype.IsImage(buf)
}

/* -------------------- Public Functions -------------------- */

func (imgFile *ImageFile) BelowMinimumDimensions(minWidth, minHeight *int) bool {
	width, height := imgFile.Dimensions()

	if (width < *minWidth) || (height < *minHeight) {
		return true
	} else {
		return false
	}
}

func (imgFile *ImageFile) GenerateHash() {
	bytes, err := imagehash.Ahash(imgFile.Image, 16)
	if err == nil {
		imgFile.Hash = hex.EncodeToString(bytes)
	}

}

func (imgFile *ImageFile) Dimensions() (width, height int) {
	size := imgFile.Image.Bounds().Size()
	width, height = size.X, size.Y

	return width, height
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

/* -------------------- Private Functions -------------------- */
