package gopic

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

//GetPixel returns RGB of the pixel at (x, y).Not Alpha
func (p *PicFile) GetPixel(x, y int) []int {
	//判断图片颜色类型
	switch p.Mode() {
	case "color.Gray":
		return GetGraypicColor(p, x, y)
	case "color.NRGBA":
		return GetRGBpicColor(p, x, y)
	case "color.CMYK":
		return GetCMYKpicColor(p, x, y)
	case "color.YCbCr":
		return GetYcbcrpicColor(p, x, y)
	default:
		return nil
	}
}

//GetGraypicColor returns Gray of the pixel at (x, y).Not Alpha
func GetGraypicColor(file *PicFile, x, y int) []int {
	col := file.img.At(x, y).(color.Gray)
	return []int{int(col.Y)}
}

//GetRGBpicColor returns RGB of the pixel at (x, y).Not Alpha
func GetRGBpicColor(file *PicFile, x, y int) []int {
	col := file.img.At(x, y).(color.NRGBA)
	return []int{int(col.R), int(col.G), int(col.B), int(col.A)}
}

//GetCMYKpicColor returns CMYK of the pixel at (x, y).Not Alpha
func GetCMYKpicColor(file *PicFile, x, y int) []int {
	col := file.img.At(x, y).(color.CMYK)
	return []int{int(col.C), int(col.M), int(col.Y), int(col.K)}
}

//GetYcbcrpicColor returns YCbCr of the pixel at (x, y).Not Alpha
func GetYcbcrpicColor(file *PicFile, x, y int) []int {
	col := file.img.At(x, y).(color.YCbCr)
	return []int{int(col.Y), int(col.Cb), int(col.Cr)}
}

//Save saves the image to file.
func (p *PicFile) Save(file string) {
	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	filetype := ImgType(file)
	switch filetype {
	case 1: //JPEG
		jpeg.Encode(f, p.img, nil)
	case 3: //PNG
		png.Encode(f, p.img)
	default:
		fmt.Println("Unsupported file type")
	}

}

//Show shows the image.
func (p *PicFile) Show() {
	//随机文件名
	a := app.New()
	w := a.NewWindow("Images")
	img := canvas.NewImageFromImage(p.img)
	w.SetContent(img)
	x, y := p.Size()
	w.Resize(fyne.NewSize(float32(x), float32(y)))
	w.ShowAndRun()
}
