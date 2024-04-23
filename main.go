package main

import (
		"net/http"
		"os"

		"github.com/go-chi/chi/v5"
)

func main() {
	router := chi,NewMux()

	// router.Get("/", )

	log.Fatal(http.ListenAndServe(os.Getenv("HTTP_LISTEN_ADDR", router)))
