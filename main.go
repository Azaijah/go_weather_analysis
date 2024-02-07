package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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

	file, err := os.Create("weather_data.csv")
	if err != nil {
		log.Fatal("Could not create file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	writer.Write([]string{"City", "Latitude", "Longitude", "Temperature", "AirPressure"})

	for _, city := range cities.Cities {
		city, err := fetchCityCoordinates(city.Name, apiKey)
		if err != nil {
			log.Fatalf("Error fetching city cords: %v", err)
		}

		weather, err := fetchWeatherData(city.Lat, city.Lon, apiKey)
		if err != nil {
			log.Fatalf("Error fetching weather data: %v", err)
		}

		// Write data for each city
		writer.Write([]string{city.Name, fmt.Sprintf("%f", city.Lat), fmt.Sprintf("%f", city.Lon), fmt.Sprintf("%f", weather.Main.Temp), fmt.Sprintf("%d", weather.Main.Pressure)})
	}
}
