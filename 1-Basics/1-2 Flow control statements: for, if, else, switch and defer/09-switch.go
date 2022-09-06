package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on")
	// 選択されたcaseだけ実行して、それ以外のcaseは実行されない
	// caseの最後にbreak Statementが自動的に提供される
	// caseは定数である必要はなく、関係する値は整数である必要がない
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n\n", os)
	}
}
