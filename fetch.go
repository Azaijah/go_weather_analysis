package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func fetchCityCoordinates(cityName, apiKey string) (City, error) {

	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", cityName, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return City{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return City{}, err

	}
	var city []City
	err = json.Unmarshal(body, &city)
	if err != nil {
		return City{}, err

	}

	return city[0], err
}

func fetchWeatherData(lat, lon float64, apiKey string) (WeatherResponse, error) {

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", lat, lon, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherResponse{}, err

	}

	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return WeatherResponse{}, err

	}

	return weather, nil
}
