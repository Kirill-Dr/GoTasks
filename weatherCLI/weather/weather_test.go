package weather_test

import (
	"strings"
	"testing"
	"weatherCLI/geo"
	"weatherCLI/weather"
)

func TestGetWeather(t *testing.T) {
	expected := "London"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3
	weatherData := weather.GetWeather(geoData, format)
	if !strings.Contains(weatherData, expected) {
		t.Errorf("Expected weather data to contain %v, but got %v", expected, weatherData)
	}
}
