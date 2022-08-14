package gopic

import (
	"golang.org/x/image/draw"
	"image"
	"image/color"
)

//ImgSplit Split the image into r,g,b channels
func (p *PicFile) ImgSplit() (r, g, b *NewPic) {
	w, h := p.Size()
	r1 := image.NewRGBA(image.Rect(0, 0, w, h))
	g1 := image.NewRGBA(image.Rect(0, 0, w, h))
	b1 := image.NewRGBA(image.Rect(0, 0, w, h))
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			red, green, blue, _ := p.img.At(i, j).RGBA()
			r1.Set(i, j, color.RGBA{uint8(red), 0, 0, 255})
			g1.Set(i, j, color.RGBA{0, uint8(green), 0, 255})
			b1.Set(i, j, color.RGBA{0, 0, uint8(blue), 255})
		}
	}
	return &NewPic{r1, w, h}, &NewPic{g1, w, h}, &NewPic{b1, w, h}
}

//Img2Array convert image to array
func (p *PicFile) Img2Array() [][][3]float32 {
	bounds := p.img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	iaa := make([][][3]float32, height)
	src_rgba := image.NewRGBA(p.img.Bounds())
	draw.Copy(src_rgba, image.Point{}, p.img, p.img.Bounds(), draw.Src, nil)

	for y := 0; y < height; y++ {
		row := make([][3]float32, width)
		for x := 0; x < width; x++ {
			idx_s := (y*width + x) * 4
			pix := src_rgba.Pix[idx_s : idx_s+4]
			row[x] = [3]float32{float32(pix[0]), float32(pix[1]), float32(pix[2])}
		}
		iaa[y] = row
	}
	return iaa
}

//Img2Array convert image to array
func (p *NewPic) Img2Array() [][][3]float32 {
	bounds := p.img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	iaa := make([][][3]float32, height)
	src_rgba := image.NewRGBA(p.img.Bounds())
	draw.Copy(src_rgba, image.Point{}, p.img, p.img.Bounds(), draw.Src, nil)

	for y := 0; y < height; y++ {
		row := make([][3]float32, width)
		for x := 0; x < width; x++ {
			idx_s := (y*width + x) * 4
			pix := src_rgba.Pix[idx_s : idx_s+4]
			row[x] = [3]float32{float32(pix[0]), float32(pix[1]), float32(pix[2])}
		}
		iaa[y] = row
	}
	return iaa
}
