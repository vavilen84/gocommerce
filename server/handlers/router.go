package handlers

import (
	"github.com/gorilla/mux"
	"github.com/vavilen84/gocommerce/constants"
	"log"
	"net/http"
	"os"
)

func MakeHandler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/sample", Sample).Methods(http.MethodPost)

	return r
}

func InitHttpServer(handler http.Handler) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)

	return &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: constants.DefaultWriteTimout,
		ReadTimeout:  constants.DefaultReadTimeout,
	}
}
