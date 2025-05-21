package main

import (
	"bufio"
	"fmt"
	"os"
)

type LineData struct {
	count   int
	inFiles map[string]bool
}

func main() {
	counts := make(map[string]*LineData)
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

	for line, data := range counts {
		if data.count > 0 {
			fmt.Printf("%d  %s : ", data.count, line)
			for name, hasLine := range data.inFiles {
				if hasLine {
					fmt.Printf("%s ", name)
				}
			}
			fmt.Printf("\n")
		}
	}
}

func countLines(f *os.File, counts map[string]*LineData) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if entry, ok := counts[input.Text()]; ok {
			entry.count++
			entry.inFiles[f.Name()] = true
		} else {
			counts[input.Text()] = &LineData{
				1,
				make(map[string]bool),
			}
		}
	}
}
