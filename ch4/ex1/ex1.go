package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func countDifferentBits(h1, h2 *[32]byte) int {
	count := 0
	for i := range h1 {
		count += int(pc[h1[i]^h2[i]])
	}
	return count
}

func main() {
	var h1, h2 [32]byte
	h1 = sha256.Sum256([]byte("XD"))
	h2 = sha256.Sum256([]byte("hehe"))
	fmt.Printf("h1: %x\nh2: %x\ndifferent bits: %d\n", h1, h2, countDifferentBits(&h1, &h2))
}
