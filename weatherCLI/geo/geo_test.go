package geo_test

import (
	"testing"
	"weatherCLI/geo"
)

func TestGetMyLocation(t *testing.T) {
	city := "London"
	expected := geo.GeoData{
		City: city,
	}

	got, err := geo.GetMyLocation(city)
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Expected city %v, but got %v", expected, got)
	}

}
