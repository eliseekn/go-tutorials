package server

import (
	"fmt"
	"http-server/routes"
	"net/http"
)

func Serve() {
	routes.Init()

	err := http.ListenAndServe(":8080", nil)

	if err == nil {
		fmt.Println("Failed to start server on port 8080")
		return
	}

	fmt.Println("Server started on port 8080...")
}
