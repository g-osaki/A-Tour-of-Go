package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	// √2 = 1.414...
	// √(-4) = 2i
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i" // 符号を反転させて、文字列iを後ろにくっつける
	}
	return fmt.Sprint(math.Sqrt(x))
}
