package server

import (
	"TemplateApi/src/service"
	"github.com/gorilla/mux"
	"net/http"
)

func makeRouter(svc service.Service) http.Handler {
	muxRouter := mux.NewRouter()

	//Heartbeat and maintenance functions
	muxRouter.HandleFunc("/heartbeat", heartbeatHandler(svc)).Methods("GET")

	//REST Call out
	muxRouter.HandleFunc("/weather", getWeatherHandler(svc)).Methods("GET")

	//Message for simple messages testing database connection
	muxRouter.HandleFunc("/message", createSimpleMessageHandler(svc)).Methods("POST")
	muxRouter.HandleFunc("/message/id/{id}", getSimpleMessageByIdHandler(svc)).Methods("GET")
	muxRouter.HandleFunc("/message/sample/{text}", getSimpleMessagesBySampleHandler(svc)).Methods("GET")
	muxRouter.HandleFunc("/message/id/{id}", deleteSimpleMessageByIdHandler(svc)).Methods("DELETE")

	//User for managing users
	muxRouter.HandleFunc("/user", createUserHandler(svc)).Methods("POST")
	muxRouter.HandleFunc("/user", getAllUsersHandler(svc)).Methods("GET")
	muxRouter.HandleFunc("/user/{id}", getUserByIDHandler(svc)).Methods("GET")
	muxRouter.HandleFunc("/user", updateUserHandler(svc)).Methods("PUT")
	muxRouter.HandleFunc("/user/{id}", deleteUserByIDHandler(svc)).Methods("DELETE")

	return muxRouter
}
