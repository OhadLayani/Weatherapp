package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getWeather(city, apiKey string) (string, error) {
	// Construct the API URL (adjust the endpoint according to the WeatherAPI documentation)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)

	// Send the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert response body to string
	jsonString := string(body)
	return jsonString, nil
}

func main() {
	// Replace with your actual API key
	apiKey := "f9a86a4bf2ce4ca0bf5142012241007"
	// Replace with the desired city
	city := "London"

	jsonData, err := getWeather(city, apiKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get weather data: %v\n", err)
		os.Exit(1)
	}

	// Write JSON data to a file
	file, err := os.Create("weather.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(jsonData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write data to file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Weather data saved to weather.json")
}
