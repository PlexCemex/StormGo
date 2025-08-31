package main

import (
	"Projects/StormGo/geo"
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "User city")
	// format := flag.Int("format", 1, "Format for output")
	flag.Parse()
	geoData, err := geo.GetMyLocation(city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(*geoData)
}

// go run . --city=""
// next 15.6