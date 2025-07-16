package main

import (
	"flag"
	"fmt"
	"weatherCLI/geo"
	"weatherCLI/weather"
)

func main() {
	fmt.Println("--- Weather CLI ---")
	city := flag.String("city", "", "User's city")
	format := flag.Int("format", 1, "Format of output")

	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}

	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
