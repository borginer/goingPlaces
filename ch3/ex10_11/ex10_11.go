package main

import (
	"bytes"
	"fmt"
	"strings"
)

func intComma(s string) string {
	if len(s) <= 3 {
		return s
	}
	var buf bytes.Buffer
	buf.WriteString(s[:len(s)%3])
	s = s[len(s)%3:]

	for len(s) >= 3 {
		buf.WriteString("," + s[:3])
		s = s[3:]
	}
	return buf.String()
}

func comma(s string) string {
	parts := strings.Split(s, ".")
	if len(parts) == 0 {
		return s
	}

	if strings.HasPrefix(parts[0], "-") {
		s = "-" + intComma(parts[0][1:])
	} else {
		s = intComma(parts[0])
	}

	if len(parts) > 1 {
		s += "." + parts[1]
	}
	return s
}

func main() {
	fmt.Println(comma("1234567891234"))
	fmt.Println(comma("21"))
	fmt.Println(comma("123"))
	fmt.Println(comma("-1123.544"))
	fmt.Println(comma("-0.524"))

}
