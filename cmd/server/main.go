package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// API endpoint
	api := r.PathPrefix("/v1").Subrouter()

	// Authority Management
	api.HandleFunc("/authorities", NotImplemented).Methods(http.MethodGet)
	api.HandleFunc("/authorities/{id}", NotImplemented).Methods(http.MethodGet)
	api.HandleFunc("/authorities", NotImplemented).Methods(http.MethodPost)
	api.HandleFunc("/authorities/{id}", NotImplemented).Methods(http.MethodDelete)
	api.HandleFunc("/authorities/{id}/sign", NotImplemented).Methods(http.MethodPost)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	srv := http.Server{
		Addr: ":" + port,
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
