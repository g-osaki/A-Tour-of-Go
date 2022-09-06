package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow { // Sliceをrangeで繰り返す場合、rangeは反復ごとに2つの変数を返す
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// 17-Range continued
	pow = make([]int, 10)

	// インデックスや値は"_"を代入することで捨てることができる
	// インデックスだけ必要であれば、2つ目の値を省略できる
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}
