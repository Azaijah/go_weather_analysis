package main

import (
	"fmt"
	"log"
)

func main() {

	config, err := loadConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	apiKey := config.API_KEY

	cities, err := loadCities("cities.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Loop over the cities and print their names and country codes
	for _, city := range cities.Cities {

		city, err := fetchCityCoordinates(city.Name, apiKey)
		if err != nil {
			log.Fatalf("Error fetching city cords: %v", err)
		}

		weather, err := fetchWeatherData(city.Lat, city.Lon, apiKey)
		if err != nil {
			log.Fatalf("Error fetching weather data: %v", err)
		}

		fmt.Printf("lon: %f, lat: %f \n", city.Lon, city.Lat)
		fmt.Printf("Air pressure: %d\nTemperature: %f\n\n", weather.Main.Pressure, weather.Main.Temp)

	}
}
