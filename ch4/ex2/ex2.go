package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var sha_type int
	flag.IntVar(&sha_type, "t", 256, "sha type to use")
	flag.Parse()

	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var s []byte
	switch sha_type {
	case 256:
		ret := sha256.Sum256(stdin)
		s = ret[:]
	case 384:
		ret := sha512.Sum384(stdin)
		s = ret[:]
	case 512:
		ret := sha512.Sum512(stdin)
		s = ret[:]
	default:
		ret := sha256.Sum256(stdin)
		s = ret[:]
	}

	fmt.Printf("sha%d: %x\n", sha_type, s)
}
