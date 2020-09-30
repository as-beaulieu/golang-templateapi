package server

import (
	"TemplateApi/src/models"
	"TemplateApi/src/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func createUserHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			respondWithJSON(w, r, http.StatusBadRequest, r.Body)
			return
		}
		defer r.Body.Close()

		result, err := svc.CreateUser(user)
		if err != nil {
			respondWithJSON(w, r, http.StatusInternalServerError, r.Body)
			return
		}

		respondWithJSON(w, r, http.StatusCreated, result)
	}
}

func getAllUsersHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := svc.GetUsers()
		if err != nil {
			respondWithJSON(w, r, http.StatusInternalServerError, r.Body)
			return
		}

		respondWithJSON(w, r, http.StatusOK, result)
	}
}

func getUserByIDHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		result, err := svc.GetUserByID(id)
		if err != nil {
			respondWithJSON(w, r, 500, nil)
		}

		respondWithJSON(w, r, http.StatusCreated, result)
	}
}

func updateUserHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			respondWithJSON(w, r, http.StatusBadRequest, r.Body)
			return
		}
		defer r.Body.Close()

		result, err := svc.UpdateUser(user)
		if err != nil {
			respondWithJSON(w, r, http.StatusInternalServerError, r.Body)
			return
		}

		respondWithJSON(w, r, http.StatusAccepted, result)
	}
}

func deleteUserByIDHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		err := svc.DeleteUser(id)
		if err != nil {
			respondWithJSON(w, r, 500, nil)
		}

		respondWithJSON(w, r, http.StatusNoContent, nil)
	}
}
