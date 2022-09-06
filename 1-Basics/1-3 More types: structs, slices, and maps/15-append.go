package main

import "fmt"

func main() {
	var s []int
	appendPrintSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	appendPrintSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	appendPrintSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	appendPrintSlice(s)

}
func appendPrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
