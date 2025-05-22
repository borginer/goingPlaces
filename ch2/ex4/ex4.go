package main

import (
	"fmt"
	"time"
)

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	sum := 0
	for i := 0; i < 64; i++ {
		sum += int(x & 1)
		x = x >> 1
	}
	return sum
}

func main() {
	start := time.Now()
	sum := 0
	for i := uint64(0); i < 50000000; i++ {
		sum += PopCount(i)
	}
	fmt.Println(sum, time.Since(start))
}
