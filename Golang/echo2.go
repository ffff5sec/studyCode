package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	for _, arg := range os.Args[1:] {
		s += s + arg

	}
	fmt.Println(s)
}
