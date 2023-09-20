package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func main() {
	reqBodyStr := "client info and api key"

	resp, err := http.Post("http://localhost:8080/jwt", "application/text", bytes.NewBuffer([]byte(reqBodyStr)))
	if err != nil {
		log.Printf("failed to get data from api %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read data from response body %v", err)
	}

	log.Printf("response: %s", body)
}
