package geo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city *string) (*GeoData, error) {
	if *city != "" {
		return &GeoData{City: *city}, nil
	}
	apapi, err := http.Get("https://ipapi.co/json")
	if err != nil {
		return nil, err
	}
	if apapi.StatusCode != 200 {
		return nil, errors.New("status code fron apapi.co is not 200")
	}
	body, err := io.ReadAll(apapi.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	err = json.Unmarshal(body,&geo)
	if err != nil {
		return nil, err
	}
	return &geo, nil
}
