package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// API endpoint
	api := router.PathPrefix("/v1").Subrouter()

	// Authority Management
	api.HandleFunc("/authorities", NotImplemented).Methods(http.MethodGet)
	api.HandleFunc("/authorities/{id}", NotImplemented).Methods(http.MethodGet)
	api.HandleFunc("/authorities", NotImplemented).Methods(http.MethodPost)
	api.HandleFunc("/authorities/{id}", NotImplemented).Methods(http.MethodDelete)

	// Certificate Management
	api.HandleFunc("/authorities/{id}/", NotImplemented).Methods(http.MethodGet)
	api.HandleFunc("/authorities/{id}/sign", NotImplemented).Methods(http.MethodPost)
	api.HandleFunc("/certificate/{serial}/revoke", NotImplemented).Methods(http.MethodPost)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	srv := http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("http.Listen %s", err)
	}
}

func Datastore(ctx context.Context) (*datastore.Client, error) {
	return datastore.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
}

func NotImplemented(w http.ResponseWriter, r *http.Request) {

}
