package main

import (
	"io"
	"log"
	"net/http"
)

// versionHandler would return the service version.
func versionHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "v1")
	if err != nil {
		http.Error(w, "Error writing the service's version", http.StatusInternalServerError)
		return
	}

}

// healthHandler would return the health of the service.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Ok")
	if err != nil {
		http.Error(w, "Error writing the service's health status", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/healthz", healthHandler)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
