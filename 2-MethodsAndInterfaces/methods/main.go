package main

import (
	"fmt"
	"math"
)

func main() {
	v := Vertex{3, 4}
	// 01
	fmt.Println(v.Abs())
	// 02
	fmt.Println(normalFuncAbs(v))

	// 03
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.anyTypeAbs())

	// 04
	v = Vertex{3, 4}
	v.Scale304(10)
	fmt.Println(v.Abs304())

	// 05
	v = Vertex{3, 4}
	Scale305(&v, 10)
	fmt.Println(Abs305(v))
}
