package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[rune]int)
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar {
			continue
		}
		counts[r]++
	}

	var letters, digits, spaces int
	for r, n := range counts {
		if unicode.IsDigit(r) {
			digits += n
		} else if unicode.IsLetter(r) {
			letters += n
		} else if unicode.IsSpace(r) {
			spaces += n
		}
	}
	formt :=
		`letters: %d
digits: %d
spaces: %d
`

	fmt.Printf(formt, letters, digits, spaces)
}
