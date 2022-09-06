package main

import "math"

func Abs305(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale305(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
