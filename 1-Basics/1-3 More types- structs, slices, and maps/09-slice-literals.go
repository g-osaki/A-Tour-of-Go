package main

import "fmt"

func main() {
	// Sliceのリテラルは長さのない配列リテラルのようなもの

	//q := [6]int{2, 3, 5, 7, 11, 13} // これは配列リテラル
	q := []int{2, 3, 5, 7, 11, 13} // これは上記と同様の配列を作成、それを参照するスライスを作成している
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
