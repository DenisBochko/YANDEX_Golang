package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	go startServer("localhost:8080")
	time.Sleep(time.Second)

	response, err := sendRequest("localhost:8080")
	if err != nil {
		log.Fatal("client failed to send request:", err)
	}

	expected := "Hello from server"
	if strings.TrimSpace(response) != expected {
		log.Fatalf("expected '%s', got '%s'", expected, response)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from server"))
}

func startServer(address string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

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

func sendRequest(url string) (string, error) {
	response, err := http.Get("http://" + url)
	if err != nil {
		return "", fmt.Errorf("Ошибка get запроса: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Ошибка чтения запроса: %w", err)
	}

	return string(body), nil
}
