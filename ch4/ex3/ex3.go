package main

import "fmt"

func reverse(ap *[8]int) {
	for i, j := 0, len(ap)-1; i < j; i, j = i+1, j-1 {
		ap[i], ap[j] = ap[j], ap[i]
	}

}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	reverse(&arr)
	fmt.Printf("%v\n", arr)
}
