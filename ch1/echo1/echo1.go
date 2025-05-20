package main

import (
	"fmt"
	"os"
)

func main() {
	var s, space string
	for i := 1; i < len(os.Args); i++ {
		s += space + os.Args[i]
		space = " "
	}

	fmt.Println(s)
}
