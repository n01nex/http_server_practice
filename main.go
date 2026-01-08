package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define the source of of the files being served by the fileserver
	fs := http.FileServer(http.Dir("."))

	// defining how the "serving" is handled
	smx := http.NewServeMux()
	smx.Handle("/app/", http.StripPrefix("/app", fs))
	smx.HandleFunc("/healthz", healthz)

	// Setting up and starting the server
	server := http.Server{
		Addr:    ":8080",
		Handler: smx,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}

// Function that just runs to confirm the health of the service that it is running
func healthz(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("OK"))

}
