package main

import (
	"encoding/json"
	"lsns/model"
	"os"
)

func LoadCityFromJSON(filePath string) ([]model.City, error) {
	data, err := os.ReadFile("./city.json")	
	if err != nil {
		return nil, err
	}
	var cities []model.City
	err = json.Unmarshal(data, &cities)
	if err != nil {
		return nil, err
	}
	return cities, nil
}