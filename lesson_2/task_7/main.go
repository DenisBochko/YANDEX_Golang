package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	startServer(":8080")
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Authorized access"))
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Authorization := r.Header.Get("Authorization")

		if Authorization != "Bearer valid_token" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func startServer(address string) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", authMiddleware(mainHandler))

	srv := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("failed to start server", err)
	}
}
