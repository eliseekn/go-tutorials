package main

import (
	"fmt"
	"http-server/server"
	"log"
)

func main() {
	port := "8080"

	err := server.Serve(port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started on port ", port, "...")
}
