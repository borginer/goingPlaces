package main

import "fmt"

func areAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	charCount := make(map[rune]int)

	for _, char := range(s1) {
		charCount[char]++
	}

	for _, char := range(s2) {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(areAnagrams("abcd", "acdb"))
	fmt.Println(areAnagrams("abcdd", "abcdf"))
}