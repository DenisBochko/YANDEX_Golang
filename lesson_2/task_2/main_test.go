package main

import (
	"strings"
	"testing"
	"time"
)

func TestHTTPClient(t *testing.T) {
	go startServer("localhost:8080")
	time.Sleep(time.Second)

	response, err := sendRequest("localhost:8080")
	if err != nil {
		t.Fatal("client failed to send request:", err)
	}

	expected := "Hello from server"
	if strings.TrimSpace(response) != expected {
		t.Fatalf("expected '%s', got '%s'", expected, response)
	}
}
