package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type Weather struct {
	Name string `json:"name"`
	Main struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
}

func loadApiKey(filepath string) (apiConfigData, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return apiConfigData{}, err
	}
	var c apiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}
	return c, nil
}

func fetchWeatherData(apiKey, city string) (Weather, error) {
	url := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiKey + "&units=metric"
	resp, err := http.Get(url)
	if err != nil {
		return Weather{}, err
	}
	defer resp.Body.Close()

	var weatherData Weather
	err = json.NewDecoder(resp.Body).Decode(&weatherData)
	if err != nil {
		return Weather{}, err
	}
	return weatherData, nil
}

func main() {
	apiConfig, err := loadApiKey(".apiConfig")
	if err != nil {
		log.Fatalf("Failed to load API key: %v", err)
	}

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.Split(r.URL.Path, "/")[2]
		data, err := fetchWeatherData(apiConfig.OpenWeatherMapApiKey, city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
