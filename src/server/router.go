package server

import (
	"TemplateApi/src/endpoint"
	"TemplateApi/src/service"
	"github.com/gorilla/mux"
	"net/http"
)

func makeRouter(svc service.Service) http.Handler {
	muxRouter := mux.NewRouter()

	//Heartbeat and maintenance functions
	muxRouter.HandleFunc("/heartbeat", endpoint.HeartbeatHandler(svc)).Methods("GET")

	//REST Call out
	muxRouter.HandleFunc("/weather", endpoint.GetWeatherHandler(svc)).Methods("GET")

	//Message for simple messages testing database connection
	muxRouter.HandleFunc("/message", endpoint.CreateSimpleMessageHandler(svc)).Methods("POST")
	muxRouter.HandleFunc("/message/id/{id}", endpoint.GetSimpleMessageByIdHandler(svc)).Methods("GET")
	muxRouter.HandleFunc("/message/sample/{text}", endpoint.GetSimpleMessagesBySampleHandler(svc)).Methods("GET")
	muxRouter.HandleFunc("/message/id/{id}", endpoint.DeleteSimpleMessageByIdHandler(svc)).Methods("DELETE")

	//User for managing users
	muxRouter.HandleFunc("/user", endpoint.CreateUserHandler(svc)).Methods("POST")
	muxRouter.HandleFunc("/user", endpoint.GetAllUsersHandler(svc)).Methods("GET")
	muxRouter.HandleFunc("/user/{id}", endpoint.GetUserByIDHandler(svc)).Methods("GET")
	muxRouter.HandleFunc("/user", endpoint.UpdateUserHandler(svc)).Methods("PUT")
	muxRouter.HandleFunc("/user/{id}", endpoint.DeleteUserByIDHandler(svc)).Methods("DELETE")

	return muxRouter
}
