package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Println("Сервер запущен на порту: 8080")
	http.ListenAndServe(":8080", nil)
}
