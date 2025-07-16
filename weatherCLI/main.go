package main

import (
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "User's city")
	format := flag.Int("format", 1, "Format of output")
	flag.Parse()
	fmt.Println(*city, *format)
}
