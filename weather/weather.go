package weather

import (
	"Projects/StormGo/geo"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func GetWeather(geo *geo.GeoData, format int) (string, error) {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		return "", nil
	}
	params := url.Values{}
	params.Add("format", strconv.Itoa(format))
	baseUrl.RawQuery = params.Encode()
	wttr, err := http.Get(baseUrl.String())
	if err != nil {
		return "", err
	}
	if wttr.StatusCode != 200 {
		return "", errors.New("status code from wttr.in is not 200")
	}
	body, err := io.ReadAll(wttr.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}