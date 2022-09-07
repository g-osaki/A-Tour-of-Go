package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// Sliceは配列への参照のようなもので、どんなデータも格納しておらず、元の配列の部分列を指し示している
	// Sliceの要素を変更すると、その元となる配列の対応する要素も変更される
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
