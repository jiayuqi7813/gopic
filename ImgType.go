package gopic

import (
	"image"
	"path/filepath"
)

// ImageType represents an image type value.
type ImageType int

type PicFile struct {
	img    image.Image
	format string
	err    error
}

type NewPic struct {
	img    *image.RGBA
	width  int
	height int
}

const (
	// UNKNOWN represents an unknow image type value.
	UNKNOWN ImageType = iota
	// JPEG represents the JPEG image type.
	JPEG
	// WEBP represents the WEBP image type.
	WEBP
	// PNG represents the PNG image type.
	PNG
	// TIFF represents the TIFF image type.
	TIFF
	// GIF represents the GIF image type.
	GIF
	// PDF represents the PDF type.
	PDF
	// SVG represents the SVG image type.
	SVG
	// MAGICK represents the libmagick compatible genetic image type.
	MAGICK
	// HEIF represents the HEIC/HEIF/HVEC image type
	HEIF
	// AVIF represents the AVIF image type.
	AVIF
)

// ImageTypes stores as pairs of image types supported and its alias names.
var ImageTypes = map[ImageType]string{
	JPEG:   "jpeg",
	PNG:    "png",
	WEBP:   "webp",
	TIFF:   "tiff",
	GIF:    "gif",
	PDF:    "pdf",
	SVG:    "svg",
	MAGICK: "magick",
	HEIF:   "heif",
	AVIF:   "avif",
}

//ImgType returns the image type of the given filename.
func ImgType(filename string) ImageType {
	ext := filepath.Ext(filename)
	for k, v := range ImageTypes {
		if ext == ".jpg" || ext == ".jpeg" {
			return 1
		}
		if v == ext[1:] {
			return k
		}

	}
	return UNKNOWN
}
