package main

import "fmt"

func reverseRunes(s []byte, n int) []rune {
	runes := []rune(string(s))

	ret := runes[len(runes)-n:]
	for i := 0; i < n; i += 1 {
		ret = append(ret, runes[i])
	}

	return ret
}

func main() {
	s := []byte("abcdםולש")
	fmt.Printf("%s\n", string(reverseRunes(s, 4)))
}
