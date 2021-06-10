package endpoint

import (
	"TemplateApi/src/models"
	"TemplateApi/src/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateUserHandler(svc service.Service) http.HandlerFunc {
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

func GetAllUsersHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := svc.GetUsers()
		if err != nil {
			respondWithJSON(w, r, http.StatusInternalServerError, r.Body)
			return
		}

		respondWithJSON(w, r, http.StatusOK, result)
	}
}

func GetUserByIDHandler(svc service.Service) http.HandlerFunc {
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

func UpdateUserHandler(svc service.Service) http.HandlerFunc {
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

func DeleteUserByIDHandler(svc service.Service) http.HandlerFunc {
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
