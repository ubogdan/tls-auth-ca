package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/gorilla/mux"
	"github.com/oklog/run"
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
	if _, err := strconv.ParseInt(port, 10, 64); err != nil {
		port = "8080"
	}

	srv := http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        router,
	}
	var g run.Group

	ctx, cancel := context.WithCancel(context.Background())
	// HTTP server
	{
		g.Add(func() error { return srv.ListenAndServe() }, func(error) { srv.Shutdown(ctx) })
	}

	// Signal Handler
	{
		g.Add(
			func() error {
				c := make(chan os.Signal, 1)
				signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
				select {
				case <-ctx.Done():
					return nil
				case <-c:
					cancel()
					log.Printf("Got to go.")
					return nil
				}
			}, func(err error) {})
	}
	_ = g.Run()

}

func Datastore(ctx context.Context) (*datastore.Client, error) {
	return datastore.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
}

func NotImplemented(w http.ResponseWriter, r *http.Request) {

}
