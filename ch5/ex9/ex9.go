package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "$foo", f("foo"))
}

func main() {
	fmt.Println(expand("hello $foo $foo",
		func(s string) string {
			return s + " world"
		}))
}
