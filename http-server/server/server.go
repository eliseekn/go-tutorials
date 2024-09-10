package server

import (
	"errors"
	"http-server/routes"
	"net/http"
)

func Serve(port string) error {
	routes.Init()

	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		return errors.New("failed to start server on port " + port)
	}

	return nil
}
