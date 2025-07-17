package weather

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"weatherCLI/geo"
)

var ErrorWrongFormat = errors.New("format must be between 1 and 4")
var ErrorParseURL = errors.New("failed to parse URL")
var ErrorHTTPGet = errors.New("failed to perform HTTP GET request")
var ErrorReadBody = errors.New("failed to read response body")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrorWrongFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		return "", ErrorParseURL
	}
	params := url.Values{}
	params.Add("format", strconv.Itoa(format))
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return "", ErrorHTTPGet
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", ErrorReadBody
	}
	return string(body), nil
}
