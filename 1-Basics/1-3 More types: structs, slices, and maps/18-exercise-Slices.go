package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	fig := make([][]uint8, dy)
	for y := range fig {
		fig[y] = make([]uint8, dx)
		for x := range fig[y] {
			fig[y][x] = uint8((y + x) / 2)
		}
	}
	return fig
}

func main() {
	pic.Show(Pic)
}
