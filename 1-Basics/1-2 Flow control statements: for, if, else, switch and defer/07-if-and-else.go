package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(
		elsepow(3, 2, 10),
		elsepow(3, 3, 20),
	)
}

func elsepow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
