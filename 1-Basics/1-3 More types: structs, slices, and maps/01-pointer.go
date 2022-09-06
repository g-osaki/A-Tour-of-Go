package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i // &は、そのオペランドへのポインタを引き出す

	fmt.Println(*p) // *は、ポインタの指す先の変数を示す。ポインタpを通じてiから値を読み出している。
	*p = 21         // ポインタpを通じてiへ値を代入
	// これは"dereferencing"または"indirecting"として知られている

	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
