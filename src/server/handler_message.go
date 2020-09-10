package server

import (
	"TemplateApi/src/models"
	"TemplateApi/src/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func createSimpleMessageHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var messageRequest models.SimpleMessageRequest

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&messageRequest); err != nil {
			respondWithJSON(w, r, http.StatusBadRequest, r.Body)
			return
		}
		defer r.Body.Close()

		id, err := uuid.NewUUID()
		if err != nil {
			fmt.Println("error making uuid for simple message", err)
			respondWithJSON(w, r, 500, r.Body)
		}
		message := models.SimpleMessage{
			ID:      id.String(),
			Message: messageRequest.Message,
		}

		messageResponse, err := svc.CreateSimpleMessage(message)
		if err != nil {
			fmt.Println("error creating simple message", err)
			respondWithJSON(w, r, 500, r.Body)
		}
		respondWithJSON(w, r, http.StatusCreated, messageResponse)
	}
}

func getSimpleMessageByIdHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func getSimpleMessagesBySampleHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func deleteSimpleMessageByIdHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
