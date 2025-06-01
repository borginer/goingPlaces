package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func delims(r rune) bool {
	dels := []rune{'-', ',', '.', '"'}
	return slices.Contains(dels, r)
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	words := make(map[string]int)

	for in.Scan() {
		for w := range strings.FieldsFuncSeq(in.Text(), delims) {
			trimed := strings.Trim(w, "\",.:?!")
			trimed = strings.ToLower(trimed)
			if len(trimed) > 0 {
				words[trimed]++
			}
		}
	}

	sorted := make([]string, 0, len(words))
	for word := range words {
		sorted = append(sorted, word)
	}
	sort.Strings(sorted)

	for _, word := range sorted {
		if words[word] >= 10 {
			fmt.Printf("%s : %d\n", word, words[word])
		}
	}
}
