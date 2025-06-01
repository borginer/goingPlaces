package main

import (
	"fmt"
)

func elimAdj(a []int) []int {
	if len(a) < 2 {
		return a
	}

	last := a[0]
	idx := 1
	for i := 1; i < len(a); i += 1 {
		if a[i] != last {
			a[idx] = a[i]
			idx += 1
		}
		last = a[i]
	}

	return a[:idx]
}

func main() {
	a := []int{1, 2, 2, 2, 3, 4, 5, 5, 6, 6, 7, 1, 1, 1, 2}
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", elimAdj(a))
}
