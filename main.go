package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/jwt")
	if err != nil {
		log.Printf("failed to get data from api %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read data from response body %v", err)
	}

	log.Printf("response: %v", body)
}
