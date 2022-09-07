package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // 評価するための簡単なStatementを書くことが可能
		return v
	}
	// return v // vはifのスコープ内だけ有効なのでここでは使えない
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10), // 9
		pow(3, 3, 20), // 20
	)
}
