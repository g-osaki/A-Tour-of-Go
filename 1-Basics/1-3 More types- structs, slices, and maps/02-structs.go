package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{2, 3}

	// 03-Struct Fields
	// structのfieldは"."を使ってアクセスする
	v.X = 3
	fmt.Println(v)

	// 04-Pointers to structs
	// structのfieldはポインタを通じてアクセス可能
	p := &v
	p.X = 1e9
	fmt.Println(v)

	// 05 Struct Literals
	fmt.Println(v1, p, v2, v3)
}
