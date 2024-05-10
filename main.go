package main

import (
	"log"
	"net/http"
	"os"
	"pictoai/handler"

	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := initEverything(); err != nil {
			log.Fatal(err)
	}

	router := chi.NewMux()

	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))

	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("application running", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func initEverything() error {
	return godotenv.Load()
}