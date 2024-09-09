package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	client, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		log.Fatal("Failed to connecto to server", err)
	}

	defer client.Body.Close()

	var res Response
	jsonErr := json.NewDecoder(client.Body).Decode(&res)

	if jsonErr != nil {
		log.Fatal("Failed to decode json", jsonErr)
	}

	fmt.Printf("status: %v\nmessage: %v\n", res.Status, res.Message)
}
