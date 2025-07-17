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
	weatherData, err := weather.GetWeather(geoData, format)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(weatherData, expected) {
		t.Errorf("Expected weather data to contain %v, but got %v", expected, weatherData)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 147},
	{name: "Zero format", format: 0},
	{name: "Negative format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expectedCity := "London"
			geoData := geo.GeoData{
				City: expectedCity,
			}
			_, err := weather.GetWeather(geoData, testCase.format)
			if err != weather.ErrorWrongFormat {
				t.Errorf("Expected error %v, but got %v", weather.ErrorWrongFormat, err)
			}
		})
	}
}
