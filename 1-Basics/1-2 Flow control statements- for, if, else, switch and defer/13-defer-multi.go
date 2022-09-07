package main

import "fmt"

func main() {
	fmt.Println("counting")

	// deferに渡した関数が複数ある場合、stackされる
	// returnする時、deferに渡した関数はLIFOの順番で実行される
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
