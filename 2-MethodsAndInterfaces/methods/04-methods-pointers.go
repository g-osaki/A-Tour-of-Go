package main

import (
	"math"
)

func (v Vertex) Abs304() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Scale304(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//func main() {
//	v := Vertex304{3, 4}
//	v.Scale304(10)
//	fmt.Println(v.Abs304())
//}
