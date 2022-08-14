package gopic

import (
	"errors"
	"image/jpeg"
	"image/png"
	"os"
)

// ImageFile Import the image and determine the type according to the suffix
//only jpg or png
func ImageFile(filename string) *PicFile {
	imgType := ImgType(filename)
	switch imgType {
	case 1: //JPEG
		img, err := FileOpen(filename)
		if err != nil {
			return nil
		}
		jpgfile, _ := jpeg.Decode(img)
		return &PicFile{jpgfile, "jpg", nil}
	case 3: //PNG
		img, err := FileOpen(filename)
		if err != nil {
			return nil
		}
		pngfile, _ := png.Decode(img)
		return &PicFile{pngfile, "png", nil}
	default:
		return &PicFile{nil, "None", errors.New("unsupported image type")}
	}
}

//FileOpen Open the file and determine the type according to the suffix
func FileOpen(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}
