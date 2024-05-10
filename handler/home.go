package handler

import (
	"fmt"
	"net/http"
	"pictoai/view/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}