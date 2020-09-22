package server

import (
	"TemplateApi/src/service"
	"github.com/gorilla/mux"
	"net/http"
)

func createUserHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func getAllUsersHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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

	}
}
func deleteUserByIDHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
