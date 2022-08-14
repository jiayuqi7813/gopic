package gopic

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image"
	"image/color"
	"image/png"
	"os"
)

//NewPicFile returns a new PicFile.
func NewPicFile(width, height int, palette color.Palette) *NewPic {
	newimg := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			newimg.Set(x, y, palette.Convert(color.RGBA{}))
		}
	}
	return &NewPic{newimg, width, height}
}

//PutPixel puts the pixel at (x, y) to the image.
func (p *NewPic) PutPixel(x, y int, palette color.Palette) {

	p.img.Set(x, y, palette.Convert(color.RGBA{}))
}

//Save saves the image to file.
func (p *NewPic) Save(file string) {
	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = png.Encode(f, p.img)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//Show shows the image.
func (p *NewPic) Show() {
	//随机文件名
	a := app.New()
	w := a.NewWindow("Images")
	img := canvas.NewImageFromImage(p.img)
	w.SetContent(img)
	w.Resize(fyne.NewSize(float32(p.width), float32(p.height)))
	w.ShowAndRun()
}
