package main

import (
	"fmt"
)

const Big = 1 << 100    // 1のビットを100箇所左にずらす
const Small = Big >> 99 // Bigのビットを99箇所右にずらす

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func needInt(x int) int {
	return x*10 + 1
}
