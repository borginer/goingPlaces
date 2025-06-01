package main

import "fmt"

func rotate(s []int, n int) []int {
	s1 := s[len(s)-n:]
	for i := range len(s) - n {
		s1 = append(s1, s[i])
	}
	return s1
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	s = rotate(s, 3)
	fmt.Printf("%v\n", s)
}
