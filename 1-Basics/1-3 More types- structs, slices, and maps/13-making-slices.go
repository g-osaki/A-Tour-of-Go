package main

import "fmt"

func main() {
	// make関数はzero化した配列を割り当てて、その配列を指すSliceを返す
	a := make([]int, 5)
	makePrintSlice("a", a)

	b := make([]int, 0, 5)
	makePrintSlice("b", b)

	c := b[:2]
	makePrintSlice("c", c)

	d := c[2:5]
	makePrintSlice("d", d)
}

func makePrintSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
