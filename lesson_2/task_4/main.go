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
	// log.Println(r.Method, r.URL)

	_, err := r.Cookie("session_id")
	if err != nil {
		newCookie := http.Cookie{
			Name:     "session_id",
			Value:    "123",
		}
		http.SetCookie(w, &newCookie)
		w.Write([]byte("Welcome!"))
		return
	}

	w.Write([]byte("Welcome back!"))
}

func startServer(address string) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)

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