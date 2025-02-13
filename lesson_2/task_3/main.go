package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	startServer(":8080")
}

func Middleware(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	w.Write([]byte("Middleware Test"))
}

func startServer(address string) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Middleware)

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
