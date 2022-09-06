package main

import (
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { // methodはレシーバ引数を関数にとる
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func normalFuncAbs(v Vertex) float64 { // 通常の関数として記述
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
