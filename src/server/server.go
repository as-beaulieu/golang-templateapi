package server

import (
	"TemplateApi/src/service"
	"fmt"
	"github.com/gorilla/handlers"
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

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	server := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        handlers.CORS(originsOk, headersOk, methodsOk)(muxRouter),
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		IdleTimeout:    0,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
