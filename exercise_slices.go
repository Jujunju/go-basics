package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	gambar := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		
		gambar[y] = make([]uint8, dx)

		for x := 0; x < dx; x++ {
			gambar[y][x] = uint8(x * y)
		}
	}

	return gambar
}

func main() {
	pic.Show(Pic)
}
