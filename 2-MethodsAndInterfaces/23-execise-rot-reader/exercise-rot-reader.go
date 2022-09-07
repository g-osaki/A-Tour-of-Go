package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		return n, err
	}

	for i, v := range b {
		switch {
		case 'A' <= v && v <= 'Z': // AからZまでのbyte
			b[i] = (v-'A'+13)%26 + 'A'
		case 'a' <= v && v <= 'z': // aからzまでのbyte
			b[i] = (v-'a'+13)%26 + 'a'
		default:
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	_, err := io.Copy(os.Stdout, &r)
	if err != nil {
		fmt.Println("failed")
	}
}
