package service

import (
	"TemplateApi/src/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func (s service) GetWeather() (*models.WeatherResponse, error) {
	apiKey := os.Getenv("WEATHERAPI")
	path := "http://api.weatherapi.com/v1/current.json?key=" + apiKey + "&q=Denver"
	response, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		var weatherResponse models.WeatherResponse
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&weatherResponse); err != nil {
			fmt.Println("error converting response to correct weather response object", err)
			return nil, err
		}
		return &weatherResponse, nil
	}

	return nil, fmt.Errorf("server error connecting with weather api")
}
