package main

import (
	"fmt"
	"os"
)

func main() {
	s, space := "", ""
	for _, arg := range os.Args[1:] {
		s += space + arg
		space = " "
	}

	fmt.Println(s)
}
