package client

import (
	"encoding/json"
	"http-client/types"
	"net/http"
)

func Get(url string) (*types.Response, error) {
	client, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer client.Body.Close()

	var res types.Response
	err = json.NewDecoder(client.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
