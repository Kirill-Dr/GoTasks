package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

func main() {
	city := flag.String("city", "", "User's city")
	format := flag.Int("format", 1, "Format of output")
	flag.Parse()
	fmt.Println(*city, *format)

	reader := strings.NewReader("Hello")
	block := make([]byte, 4)
	for {
		if _, err := reader.Read(block); err == io.EOF {
			break
		}
		fmt.Printf("%q\n", block)
	}
}
