package handlers

import (
	"net/http"
)

func Readiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
