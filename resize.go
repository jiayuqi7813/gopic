package gopic

import "github.com/nfnt/resize"

//Resize resizes the image to the specified width and height.
func (p *PicFile) Resize(width, height int) *PicFile {
	m := resize.Resize(uint(width), uint(height), p.img, resize.Lanczos3)
	return &PicFile{m, p.format, nil}
}

func (p *NewPic) Resize(width, height int) *PicFile {
	m := resize.Resize(uint(width), uint(height), p.img, resize.Lanczos3)
	return &PicFile{m, "png", nil}
}
