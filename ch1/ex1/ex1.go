package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	name_split := strings.Split(os.Args[0], "/")
	prog_name := name_split[len(name_split) - 1]
	fmt.Println(prog_name + ": " + strings.Join(os.Args[1:], " "))
}