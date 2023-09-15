package handlers

import "net/http"

// Health returns a HTTP 200 status code indicating the service is alive
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
