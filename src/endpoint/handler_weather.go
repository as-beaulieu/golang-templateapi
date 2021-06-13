package endpoint

import (
	"TemplateApi/src/service"
	"net/http"
)

func GetWeatherHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		weatherResponse, err := svc.GetWeather()
		if err != nil {
			respondWithJSON(w, r, 500, r.Body)
			return
		}
		respondWithJSON(w, r, http.StatusOK, weatherResponse)
	}
}
