package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	response, err := sendRequest("http://localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(response)

}

func startServer(address string) {

}

func sendRequest(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Ошибка get запроса: %e", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Ошибка чтения запроса: %e", err)
	}
	return string(body), nil
}
