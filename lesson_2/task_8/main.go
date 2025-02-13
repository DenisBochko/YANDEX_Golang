package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	startServer(":8080")
}

func languageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("lang")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "lang",
			Value: "en",
		}
	}

	switch cookie.Value {
	case "en":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello!\n"))
		return
	case "ru":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Привет!\n"))
		return
	default:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello!\n"))
		return
	}

}

// func cookieMiddleware(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		_, err := r.Cookie("lang")
// 		if err != nil {
// 			switch {
// 			case errors.Is(err, http.ErrNoCookie): // если кука не найдена, то редиректим
// 				http.SetCookie(w, &http.Cookie{
// 					Name:  "lang",
// 					Value: "en",
// 				})
// 			default:
// 				log.Println(err)
// 				http.Error(w, "server error", http.StatusInternalServerError)
// 			}
// 		}
// 		next(w, r)
// 	}
// }

func startServer(address string) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", languageHandler)

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
