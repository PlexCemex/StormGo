package geo_test

import (
	"Projects/StormGo/geo"
	"testing"
)

// go test .\geo\geo_test.go
func TestGetMyLocation(t *testing.T) {
	// Arrange
	city := "London"
	expected := geo.GeoData {
		City: "London",
	}
	// Act
	got, err := geo.GetMyLocation(&city)
	// Assert
	if err != nil {
		t.Errorf("Получина ошибка тестирования при передаче города: %v", err)
	}
	if *got != expected {
		t.Errorf("Ожидалось %v, получино %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T){
	city := "NoLondon"
	_, err := geo.GetMyLocation(&city)
	if err != geo.ErrNoCity {
		t.Errorf("Ожидалось %v, получино %v", geo.ErrNoCity, err)
	}
}