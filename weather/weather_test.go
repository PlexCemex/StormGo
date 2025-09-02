package weather_test

import (
	"Projects/StormGo/geo"
	"Projects/StormGo/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	expected := "London"
	format := 4
	geo := geo.GeoData{
		City: expected,
	}

	gotWeather, err := weather.GetWeather(&geo, format)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(gotWeather, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, gotWeather)
	}
}

var testCasesWrongFormat = []struct {
	name   string
	format int
}{
	{"Minus format", -1},
	{"Zero format", 0},
	{"Big format", 100},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCasesWrongFormat {
		expected := "London"
		geo := geo.GeoData{
			City: expected,
		}
		_, err := weather.GetWeather(&geo, tc.format)
		if err != weather.ErrWrongFormat {
			t.Errorf("Ожидалось %v, получено %v", weather.ErrWrongFormat, err)
		}
	}
}
