package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) > 0 {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Printf("dup2: %v\n", err)
				continue
			} else {
				countLines(f, counts)
			}
		}
	} else {
		countLines(os.Stdin, counts)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d  %s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
