package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy) //make slice length of dy
	
	for y := 0; y < dy; y++{
		img[y] = make([]uint8, dx) //manipulate []uint8
		
		for x := 0;x < dx;x++{
			img[y][x] = uint8((y^x)*(y^x)) //change the equation to get diff greyscale pics
		}
	}
	return img
}

func main() {
	pic.Show(Pic)
}