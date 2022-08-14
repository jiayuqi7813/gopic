package gopic

import (
	"fmt"
	"image/color"
)

//Size returns the size of the image.
func (p *PicFile) Size() (int, int) {
	width := p.img.Bounds().Max.X
	height := p.img.Bounds().Max.Y
	return width, height
}

//Format returns the format of the image.
func (p *PicFile) Format() string {
	return p.format
}

//Mode returns the mode of the image.
func (p *PicFile) Mode() string {
	//col := color.RGBA{}
	//colmode := p.img.ColorModel().Convert(col)
	//return fmt.Sprintf("%T", colmode)
	//输出图片色彩模式
	c := color.ModelFunc(p.img.ColorModel().Convert)
	return fmt.Sprintf("%T", c.Convert(color.RGBA{}))
}
