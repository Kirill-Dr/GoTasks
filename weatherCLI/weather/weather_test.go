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

func TestGetWeatherWrongFormat(t *testing.T) {
	expectedCity := "London"
	geoData := geo.GeoData{
		City: expectedCity,
	}
	format := 100
	_, err := weather.GetWeather(geoData, format)
	if err != weather.ErrorWrongFormat {
		t.Errorf("Expected error %v, but got %v", weather.ErrorWrongFormat, err)
	}
}
