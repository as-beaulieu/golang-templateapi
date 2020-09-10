package server

import (
	"TemplateApi/src/service"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func RunHttpServer(svc service.Service) error {
	muxRouter := makeRouter(svc)
	httpAddr := os.Getenv("PORT")
	if httpAddr == "" {
		fmt.Println("port is empty from env variables, setting to default")
		httpAddr = "8080"
	}
	fmt.Println("Listening on ", httpAddr)

	server := &http.Server{
		Addr:	":" + httpAddr,
		Handler: muxRouter,
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout: 0,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
