package main

import (
	"image"
	"image/color"
	"image/png"
	"os"        
)

type Image struct{}

func (i Image) Bounds() image.Rectangle { return image.Rect(0, 0, 256, 256) }
func (i Image) ColorModel() color.Model { return color.RGBAModel }
func (i Image) At(x, y int) color.Color {
	v := uint8(x ^ y) 
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}

	fileGambar, err := os.Create("lukisan_jujun.png")
	if err != nil {
		panic(err)
	}
	defer fileGambar.Close()

	png.Encode(fileGambar, m)
	
}
