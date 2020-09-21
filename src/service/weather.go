package service

import (
	"TemplateApi/src/models"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func (s service) GetWeather() (*models.WeatherResponse, error) {
	logger := s.logger.Named("s.GetWeather")

	logger.Info("Calling for weather report")

	apiKey := os.Getenv("WEATHERAPI")
	path := "http://api.weatherapi.com/v1/current.json?key=" + apiKey + "&q=Denver"
	response, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		logger.Info("Call for weather report returned status OK")
		var weatherResponse models.WeatherResponse
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&weatherResponse); err != nil {
			fmt.Println("error converting response to correct weather response object", err)
			return nil, err
		}
		return &weatherResponse, nil
	}

	logger.Error("error calling for weather report", zap.Int("status_code", response.StatusCode))
	return nil, fmt.Errorf("server error connecting with weather api")
}
