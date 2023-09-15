package handlers

import (
	"net/http"
)

// Root returns a HTTP 200 status code
func Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
