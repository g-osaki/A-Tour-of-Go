package main

import "fmt"

type Vertex20 struct {
	Lat, Long float64
}

var mm = map[string]Vertex20{
	"Bell Labs": Vertex20{40.68433, -74.39967},
	"Google":    Vertex20{37.42202, -122.08408},
}

// 21-map-literals-continued
// mmと同じ
var nn = map[string]Vertex20{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Printf("mm : %v\nnn : %v", mm, nn)
}
