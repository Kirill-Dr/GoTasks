package main

import (
	"flag"
	"fmt"
	"weatherCLI/geo"
)

func main() {
	city := flag.String("city", "", "User's city")
	// format := flag.Int("format", 1, "Format of output")
	flag.Parse()
	fmt.Println(*city)

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)
}
