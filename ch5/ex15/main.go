package main

import (
	"fmt"
)

func selectInt(c func(int, int) bool, vals ...int) int {
	ret := vals[0]
	for _, num := range vals[1:] {
		if c(ret, num) {
			ret = num
		}
	}
	return ret
}

func min(a int, more ...int) int {
	c := func(a, b int) bool {
		return a > b
	}
	return selectInt(c, append(more, a)...)
}

func max(a int, more ...int) int {
	c := func(a, b int) bool {
		return a < b
	}
	return selectInt(c, append(more, a)...)
}

func main() {
	fmt.Println(min(5, 4, 6, 3, 2))
	fmt.Println(max(5, 4, 6, 3, 2))
}
