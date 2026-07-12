// package main

// import (
// 	"fmt"
// 	"image"
// 	"image/color"
// 	"image/png"
// 	"os"

// 	"golang.org/x/tour/pic"
// )

// type Image struct{}

// func main() {

// 	m := Image{}

// 	pic.ShowImage(m)
	
// 	 file, err := os.Create("by_j.png")  
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer file.Close()

// 	png.Encode(file, m)

// }

// func (i Image) Bounds() image.Rectangle {
// 	return image.Rect(0, 0, 155, 155)
// }

// func (i Image) ColorModel() color.Model {
// 	return color.RGBAModel
// }

// func (i Image) At(x, y int) color.Color {

// 	v := uint8(x * y * 8 + 100 - 55)

// 	return color.RGBA{v, v, 230, 133}

// }