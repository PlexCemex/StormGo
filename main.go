package main

import (
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "User city")
	format := flag.Int("format", 1, "Format for output")

	flag.Parse()

	fmt.Println(*city, *format)
}

// go run . --city=""