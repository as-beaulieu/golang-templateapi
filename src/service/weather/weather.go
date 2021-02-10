package weather

import (
	"TemplateApi/src/models"
	"TemplateApi/src/service"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
)

type local_service struct {
	*service.TemplateService
}

type WeatherReporter interface {
	GetWeather() (*models.WeatherResponse, error)
}

func (s *local_service) GetWeather() (*models.WeatherResponse, error) {
	logger := s.Logger.Named("s.GetWeather")

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
