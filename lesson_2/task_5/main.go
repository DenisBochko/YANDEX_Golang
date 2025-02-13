package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	startServer(":8080")
}

func myMiddleware(w http.ResponseWriter, r *http.Request) {
	// log.Println(r.Method, r.URL)

	_, err := r.Cookie("user_id")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie): // если кука не найдена, то редиректим
			http.Redirect(w, r, "/login", 301)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	w.Write([]byte("Access granted"))
}

func login(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "user_id",
		Value: "123",
	}
	http.SetCookie(w, &cookie)
	w.Write([]byte("Please log in"))
}

func startServer(address string) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", myMiddleware)
	mux.HandleFunc("/login", login)

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
