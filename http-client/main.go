package main

import (
	"fmt"
	"http-client/client"
	"log"
)

func main() {
	res, err := client.Get("http://127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("status: %v\nmessage: %v\n", res.Status, res.Message)
}
