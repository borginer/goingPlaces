package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
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
