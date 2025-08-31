package main

import (
	"Projects/StormGo/geo"
	"Projects/StormGo/weather"
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "User city")
	format := flag.Int("format", 1, "Format for output")
	flag.Parse()
	geoData, err := geo.GetMyLocation(city)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	weatherData, err := weather.GetWeather(geoData, *format)
	if err != nil {
		fmt.Println(err.Error())
	}
	if *format == 3 || *format == 4 {
		fmt.Print( weatherData)
		return
	}
	fmt.Print(geoData.City, ": ", weatherData)
}

// Example: go run . --city="" --format=4
// next 16.1