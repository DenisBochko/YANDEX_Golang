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
	w.Write([]byte("Access granted"))
}

func ipBlockerMiddleware(blockedIP string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Real-IP")

		if ip == blockedIP {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

func startServer(address string) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", ipBlockerMiddleware("192.168.0.1", mainHandler))

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

// func getIP(r *http.Request) string {
// 	ip, _, err := net.SplitHostPort(r.RemoteAddr)
// 	if err != nil {
// 		return ""
// 	}
// 	return ip
// }
