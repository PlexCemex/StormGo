package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type cityPopulationResponce struct {
	Error bool `json:"error"`
}

var ErrNoCity = errors.New("entered city does not exist")
var ErrNot200 = errors.New("status code from apapi.co is not 200")
func GetMyLocation(city *string) (*GeoData, error) {
	if *city != "" {
		if ! chekcCity(*city) {
			return nil, ErrNoCity
		}
		return &GeoData{City: *city}, nil
	}
	apapi, err := http.Get("https://ipapi.co/json")
	if err != nil {
		return nil, err
	}
	if apapi.StatusCode != 200 {
		return nil, ErrNot200
	}
	body, err := io.ReadAll(apapi.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	err = json.Unmarshal(body, &geo)
	if err != nil {
		return nil, err
	}
	return &geo, nil
}

func chekcCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	countriesnow, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}
	if countriesnow.StatusCode != 200 {
		return false
	}
	body, err := io.ReadAll(countriesnow.Body)
	if err != nil {
		return false
	}
	var PopulationResponce cityPopulationResponce
	err = json.Unmarshal(body, &PopulationResponce)
	if err != nil {
		return false
	}
	return ! PopulationResponce.Error
}
