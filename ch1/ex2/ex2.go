package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println("Arg" + strconv.Itoa(i+1) + ": " + arg)
	}
}
