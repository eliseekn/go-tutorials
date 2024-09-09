package routes

import (
	"http-server/handlers"
	"net/http"
)

func Init() {
	http.HandleFunc("/", handlers.Index)
}
