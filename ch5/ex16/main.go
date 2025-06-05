package main

import "fmt"

func Join2(sep string, fields ...string) string {
	if len(fields) == 0 {
		return ""
	}

	ret := fields[0]
	for _, s := range fields[1:] {
		ret += sep + s
	}
	return ret
}

func main() {
	fmt.Printf("%s\n", Join2(" nice ", "a", "b", "c", "what do you mean by good morning"))
}
