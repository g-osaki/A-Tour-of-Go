package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println("-----------------")
	fmt.Println("Result    : ", Sqrt(2))
	fmt.Println("math.Sqrt : ", math.Sqrt(2))
}
